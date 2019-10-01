package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	game_pb "github.com/ericfengchao/treasure-hunting/protos"
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
	// register with tracker, confirm own identity
	playerId := os.Args[3]

	grpcListener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen for grpc: %v", err)
	}

	gameSvc := service.NewGameSvc(service.PrimaryNode, playerId, gridSize, treasureAmount, nil)
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
