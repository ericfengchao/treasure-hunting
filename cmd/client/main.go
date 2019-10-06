package main

import (
	"context"
	"flag"
	game_pb "github.com/ericfengchao/treasure-hunting/protos"
	"google.golang.org/grpc"
	"log"
)

var player = flag.String("player", "FC", "player's id")
var move = flag.Int("move", 0, "the move you want for the player")

func main() {
	flag.Parse()
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed to connect to server", err)
	}
	defer conn.Close()
	client := game_pb.NewGameServiceClient(conn)
	resp, err := client.MovePlayer(context.Background(), &game_pb.MoveRequest{
		Id:   *player,
		Move: int32(*move),
	})
	log.Println(resp, err)
}
