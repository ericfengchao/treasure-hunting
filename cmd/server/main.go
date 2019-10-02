package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	game_pb "github.com/ericfengchao/treasure-hunting/protos"
	tracker_pb "github.com/ericfengchao/treasure-hunting/protos/tracker"
	"github.com/ericfengchao/treasure-hunting/service"
	"google.golang.org/grpc"
)

var gridSize = 3
var treasureAmount = 5

func main() {
	if len(os.Args) < 4 {
		panic("insufficient params to start the game")
	}
	//trackerHost := os.Args[1]
	//trackerPort := os.Args[2]
	//trackerAddress := trakerHost + ":" + trackerPort
	trackerAddress := "localhost:51000"
	// Register with tracker, confirm own identity
	playerId := os.Args[3]
	role, N, K, port, registry, err := Register(trackerAddress, playerId)
	if err != nil {
		log.Fatalf("Failed to register %v", err)
	}
	assign_address := "localhost:" + string(port)
	grpcListener, err := net.Listen("tcp", assign_address)
	if err != nil {
		log.Fatalf("Failed to listen for grpc: %v", err)
	}
	gameSvc := service.NewGameSvc(role, playerId, N, K, registry)
	svr := grpc.NewServer()
	game_pb.RegisterGameServiceServer(svr, gameSvc)

	go func() {
		if err := svr.Serve(grpcListener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	fmt.Println("GAME STARTS NOW")

	// http
	http.Handle("/", gameSvc)
	if err := http.ListenAndServe("localhost:50052", nil); err != nil {
		log.Fatalf("failed to start http server: %v", err)
	}
}

func Register(trackerAddress string, playerId string) (service.Role, int, int, int, *game_pb.Registry, error) {
	// First request tracker, to make sure which role I am
	var role service.Role
	conn, err := grpc.Dial(trackerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed to connect to server", err)
	}
	defer conn.Close()
	client := tracker_pb.NewTrackerServiceClient(conn)
	resp, err := client.Register(context.Background(), &tracker_pb.RegisterRequest{
		PlayerId: playerId,
	})

	if resp.Status == tracker_pb.RegisterResponse_REGISTERED {
		return service.PrimaryNode, 0, 0, 0, nil, errors.New("Registered")
	}
	registry := ConvertRegistryType(resp.Registry)
	for k, v := range resp.Registry.PlayerList {
		if v.PlayerId == playerId {
			if k == 0 {
				role = service.PrimaryNode
			}
			if k == 1 {
				role = service.BackupNode
			} else {
				role = service.PlayerNode
			}
		}
	}

	N := int(resp.N)
	K := int(resp.K)
	port := int(resp.AssignedPort)

	return role, N, K, port, registry, nil
}

func ConvertRegistryType(registry *tracker_pb.Registry) *game_pb.Registry {
	version := registry.Version
	playerlist := registry.PlayerList
	var gameRegistry []*game_pb.Registry_Player
	for k, v := range playerlist {
		player := &game_pb.Registry_Player{
			Id:        v.PlayerId,
			JoinOrder: int32(k),
		}
		gameRegistry = append(gameRegistry, player)
	}
	return &game_pb.Registry{
		Version: version,
		Players: gameRegistry,
	}
}
