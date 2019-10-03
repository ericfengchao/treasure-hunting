package main

import (
	"context"
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
	// register with tracker, confirm own identity
	playerId := os.Args[3]
	grpcListener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen for grpc: %v", err)
	}
	role := getMyRole(trackerAddress, playerId)
	gameSvc := service.NewGameSvc(role, playerId, gridSize, treasureAmount, nil)
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

func getMyRole(trackerAddress string, playerId string) service.Role {
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

	for k, v := range resp.Registry.PlayerList {
		if v.PlayerId == playerId {
			if k == 0 {
				role = service.PrimaryNode
			}
			if k == 1 {
				role = service.BackupNode
			}
		}
	}
	return role
}
