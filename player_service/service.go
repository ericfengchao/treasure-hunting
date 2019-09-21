package player_service

import (
	"context"
	"fmt"
	game_pb "github.com/ericfengchao/treasure-hunting/protos"
	"github.com/ericfengchao/treasure-hunting/service"
	"github.com/ericfengchao/treasure-hunting/service/models"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"
)

type playerSvc struct {
	tracker     game_pb.TrackerServiceClient
	trackerConn *grpc.ClientConn
	registry    *game_pb.Registry

	gridSize, treasureAmount int

	shutdown chan struct{}

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
}

func (p *playerSvc) Close() {
	if p.trackerConn != nil {
		p.trackerConn.Close()
	}
}

func (p *playerSvc) deriveRole() models.Role {
	if p.getPrimaryServer().GetPlayerId() == p.id {
		return models.PrimaryNode
	} else if p.getBackupServer().GetPlayerId() == p.id {
		return models.BackupNode
	} else {
		return models.PlayerNode
	}
}

func (p *playerSvc) getPrimaryServer() *game_pb.Player {
	if len(p.registry.GetPlayerList()) > 0 {
		return p.registry.GetPlayerList()[0]
	}
	return nil
}

func (p *playerSvc) getBackupServer() *game_pb.Player {
	if len(p.registry.GetPlayerList()) > 1 {
		return p.registry.GetPlayerList()[1]
	}
	return nil
}

func (p *playerSvc) getHeartbeatPlayers() (prev, next *game_pb.Player) {
	if len(p.registry.GetPlayerList()) <= 1 {
		return nil, nil
	}
	for i, node := range p.registry.GetPlayerList() {
		if node.PlayerId == p.id {
			next = p.registry.GetPlayerList()[(i+1)%len(p.registry.GetPlayerList())]
			break
		}
		prev = node
	}
	return
}

func connectToPlayer(player *game_pb.Player) game_pb.GameServiceClient {
	if player == nil {
		return nil
	}
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", player.GetIp(), player.GetPort()),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Println("err connecting to player", player, err)
		return nil
	}
	return game_pb.NewGameServiceClient(conn)
}

func (p *playerSvc) refreshHeartbeatNodes() {
	prev, next := p.getHeartbeatPlayers()
	if prev.GetPlayerId() != p.prevNode.GetPlayerId() {
		// previous neighbour is changed now, connect to the new one
		p.prevNode = prev
		p.prevNodeClient = connectToPlayer(prev)
	}
	if next.GetPlayerId() != p.prevNode.GetPlayerId() {
		// next neighbour is changed now, connect to the new one
		p.nextNode = next
		p.nextNodeClient = connectToPlayer(next)
	}
}

func (p *playerSvc) refreshPrimaryNode() {
	primary := p.getPrimaryServer()
	if primary.GetPlayerId() != p.gamePrimary.GetPlayerId() {
		p.gamePrimary = primary
		p.gamePrimaryClient = connectToPlayer(primary)
	}
}

func (p *playerSvc) StartHeartbeat() {
	t := time.NewTicker(time.Millisecond * 500)
	ctx := context.Background()
	for {
		select {
		case <-t.C:
			p.refreshHeartbeatNodes()
			if p.prevNodeClient != nil {
				p.prevNodeClient.Heartbeat(ctx, &game_pb.HeartbeatRequest{
					PlayerId: p.id,
					//Registry: p.registry,
				})
			}
			if p.nextNodeClient != nil {
				p.nextNodeClient.Heartbeat(ctx, &game_pb.HeartbeatRequest{
					PlayerId: p.id,
					//Registry: p.registry,
				})
			}
		case _, open := <-p.shutdown:
			if !open {
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
	p.gameSvc = service.NewGameSvc(p.deriveRole(), p.id, p.gridSize, p.treasureAmount, p.registry)
	svr := grpc.NewServer()
	game_pb.RegisterGameServiceServer(svr, p.gameSvc)

	go func() {
		if err := svr.Serve(grpcListener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	log.Println("player grpc starts listening now")

	go p.StartHeartbeat()

	log.Println("player heartbeat starts now")

	// http
	http.Handle("/", p.gameSvc)
	if err := http.ListenAndServe(fmt.Sprintf("localhost:%d", p.assignedPort+1), nil); err != nil {
		log.Fatalf("failed to start http server: %v", err)
	}
}

func (p *playerSvc) Start(closing chan<- struct{}) {
	defer func() {
		closing <- struct{}{}
	}()
	rand.Seed(time.Now().Unix())
	ctx := context.Background()
	// start own server
	p.refreshPrimaryNode()
	t := time.NewTicker(time.Second * 2)
	// todo mocking only
	i := 0
	for {
		select {
		case <-t.C:
			if i >= 15 {
				close(p.shutdown)
				break
			}
			row, col := rand.Intn(p.gridSize), rand.Intn(p.gridSize)
			resp, err := p.gamePrimaryClient.TakeSlot(ctx, &game_pb.TakeSlotRequest{
				Id: p.id,
				MoveToCoordinate: &game_pb.Coordinate{
					Row: int32(i % p.gridSize),
					Col: int32(i % p.gridSize),
				},
			})
			if err != nil {
				log.Println(err)
			} else {
				log.Println(fmt.Sprintf("player-%s row-%d col-%d", p.id, row, col), resp.Status.String())
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
	}
}
