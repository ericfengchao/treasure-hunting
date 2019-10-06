package player_service

import (
	"bufio"
	"context"
	"fmt"
	"github.com/ericfengchao/treasure-hunting/service/models"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	game_pb "github.com/ericfengchao/treasure-hunting/protos"
	"github.com/ericfengchao/treasure-hunting/service"
	"google.golang.org/grpc"
)

type playerSvc struct {
	tracker     game_pb.TrackerServiceClient
	trackerConn *grpc.ClientConn
	registry    *game_pb.Registry

	gridSize, treasureAmount int

	shutdown chan struct{}
	wg       *sync.WaitGroup

	// primary server
	gamePrimaryClient game_pb.GameServiceClient
	gamePrimary       *game_pb.Player

	// prev heartbeat
	prevNodeClient game_pb.GameServiceClient
	prevNode       *game_pb.Player

	// next heartbeat
	nextNodeClient game_pb.GameServiceClient
	nextNode       *game_pb.Player

	assignedPort int
	id           string

	playerStates []*game_pb.PlayerState
	gameSvc      service.GameService
	listener     net.Listener
}

func (p *playerSvc) Close() {
	if p.trackerConn != nil {
		p.trackerConn.Close()
	}
}

func (p *playerSvc) getHeartbeatPlayers() (prev, next *game_pb.Player) {
	if len(p.registry.GetPlayerList()) <= 1 {
		return nil, nil
	}
	numOfPlayer := len(p.registry.GetPlayerList())
	for i, node := range p.registry.GetPlayerList() {
		if node.PlayerId == p.id {
			next = p.registry.GetPlayerList()[(i+1)%numOfPlayer]
			prev = p.registry.GetPlayerList()[(i-1+numOfPlayer)%numOfPlayer]
			break
		}
		prev = node
	}
	return
}

func (p *playerSvc) refreshHeartbeatNodes() {
	p.registry = p.gameSvc.GetLocalRegistry()
	prev, next := p.getHeartbeatPlayers()
	if prev.GetPlayerId() != p.prevNode.GetPlayerId() {
		// previous neighbour is changed now, connect to the new one
		p.prevNode = prev
		p.prevNodeClient = service.ConnectToPlayer(prev)
	}
	if next.GetPlayerId() != p.prevNode.GetPlayerId() {
		// next neighbour is changed now, connect to the new one
		p.nextNode = next
		p.nextNodeClient = service.ConnectToPlayer(next)
	}
}

func (p *playerSvc) refreshPrimaryNode() {
	primary := service.GetPrimaryServer(p.registry)
	if p.gamePrimary == nil || primary.GetPlayerId() != p.gamePrimary.GetPlayerId() {
		p.gamePrimary = primary
		p.gamePrimaryClient = service.ConnectToPlayer(primary)
	}
}

// contact tracker to report neighbour missing
func (p *playerSvc) reportMissingNode(ctx context.Context, playerId string) {
	resp, err := p.tracker.ReportMissing(ctx, &game_pb.Missing{PlayerId: playerId})
	if err != nil {
		log.Println("report missing failed", playerId, err)
		// best effort
		return
	}
	p.registry = resp.GetRegistry()
	p.gameSvc.UpdateLocalRegistry(p.registry)
}

func (p *playerSvc) StartHeartbeat() {
	p.wg.Add(1)
	defer p.wg.Done()
	t := time.NewTicker(time.Millisecond * 500)
	ctx := context.Background()
	for {
		select {
		case <-t.C:
			p.refreshHeartbeatNodes()
			if p.prevNodeClient != nil {
				_, err := p.prevNodeClient.Heartbeat(ctx, &game_pb.HeartbeatRequest{
					PlayerId: p.id,
					Registry: p.registry,
				})
				if err != nil {
					// player is missing
					p.reportMissingNode(ctx, p.prevNode.PlayerId)
				}
			}
			if p.nextNodeClient != nil {
				_, err := p.nextNodeClient.Heartbeat(ctx, &game_pb.HeartbeatRequest{
					PlayerId: p.id,
					Registry: p.registry,
				})
				if err != nil {
					// player is missing
					p.reportMissingNode(ctx, p.nextNode.PlayerId)
				}
			}
		case _, open := <-p.shutdown:
			if !open {
				p.listener.Close()
				log.Println("receive shutting down signal")
				return
			}
		}
	}
}

func (p *playerSvc) StartServing() {
	grpcListener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", p.assignedPort))
	if err != nil {
		log.Fatalf("Failed to listen for grpc: %v", err)
	}
	p.listener = grpcListener
	p.gameSvc = service.NewGameSvc(p.id, p.gridSize, p.treasureAmount, p.registry)
	svr := grpc.NewServer()
	game_pb.RegisterGameServiceServer(svr, p.gameSvc)

	go func() {
		if err := svr.Serve(grpcListener); err != nil {
			log.Printf("failed to serve: %v", err)
			return
		}
	}()

	log.Println("player grpc starts listening now")

	go p.StartHeartbeat()
	log.Println("player heartbeat starts now")
	// http
	http.Handle("/", p.gameSvc)
	if err := http.ListenAndServe(fmt.Sprintf("localhost:%d", p.assignedPort+1), nil); err != nil {
		log.Printf("failed to start http server: %v", err)
		return
	}
}

func (p *playerSvc) KeyboardListen(closing chan<- struct{}) {
	defer func() {
		closing <- struct{}{}
	}()
	ctx := context.Background()
	p.refreshPrimaryNode()
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		input := reader.Text()
		move, err := service.ParseDirection(input)
		if err != nil {
			log.Printf("fail to parse the input, err: %s", err.Error())
			return
		}
		switch move {
		case models.Exit:
			close(p.shutdown)
			p.wg.Wait()
			p.Close()
			log.Println("receive shutting down signal")
			return
		case models.Up, models.Right, models.Down, models.Left:
			resp, err := p.gamePrimaryClient.MovePlayer(ctx, &game_pb.MoveRequest{
				Id:   p.id,
				Move: int32(move),
			})
			if err != nil {
				log.Println(err)
			} else {
				log.Println(fmt.Sprintf("player-%s move %d done", p.id, move), resp.Status.String())
				p.playerStates = resp.GetPlayerStates()
			}
		}
	}
}

func (p *playerSvc) Start(closing chan<- struct{}) {
	defer func() {
		closing <- struct{}{}
	}()
	rand.Seed(time.Now().Unix())
	ctx := context.Background()
	t := time.NewTicker(time.Second * 2)
	// todo mocking only
	i := 0
	for {
		p.refreshPrimaryNode()
		select {
		case <-t.C:
			if p.id == "11" && i >= 3 {
				close(p.shutdown)
				p.wg.Wait()
				return
			} else if i >= 10 {
				close(p.shutdown)
				p.wg.Wait()
				return
			}
			move := int32(rand.Intn(5))
			resp, err := p.gamePrimaryClient.MovePlayer(ctx, &game_pb.MoveRequest{
				Id:   p.id,
				Move: move,
			})
			if err != nil {
				log.Println(err)
			} else {
				log.Println(fmt.Sprintf("player-%s move-%d", p.id, move), resp.Status.String())
				p.playerStates = resp.GetPlayerStates()
			}
			i++
		}
	}
}

func NewPlayerSvc(trackerHost, trackerPort string, id string) *playerSvc {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", trackerHost, trackerPort), grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed to connect to tracker", err)
	}
	trackerClient := game_pb.NewTrackerServiceClient(conn)

	// register with tracker, confirm own identity
	resp, err := trackerClient.Register(context.Background(), &game_pb.RegisterRequest{PlayerId: id})
	if err != nil {
		log.Fatal("failed to bootstrap with tracker", err)
	} else {
		log.Println("registration done!", resp)
	}

	return &playerSvc{
		tracker:        trackerClient,
		trackerConn:    conn,
		registry:       resp.GetRegistry(),
		gamePrimary:    nil, // populate
		assignedPort:   int(resp.GetAssignedPort()),
		id:             id,
		gridSize:       int(resp.GetN()),
		treasureAmount: int(resp.GetK()),
		shutdown:       make(chan struct{}, 0),
		wg:             &sync.WaitGroup{},
	}
}
