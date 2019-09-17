package main

import (
	"context"
	"flag"
	game_pb "github.com/ericfengchao/treasure-hunting/service/protos"
	"google.golang.org/grpc"
	"log"
)

var player = flag.String("player", "FC", "player's id")
var row = flag.Int("row", 0, "the row of the coordinates that you want to move to")
var col = flag.Int("col", 0, "the col of the coordinates that you want to move to")

func main() {
	flag.Parse()
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed to connect to server", err)
	}
	defer conn.Close()
	client := game_pb.NewGameServiceClient(conn)
	resp, err := client.TakeSlot(context.Background(), &game_pb.TakeSlotRequest{
		Id: *player,
		MoveToCoordinate: &game_pb.Coordinate{
			Row: int32(*row),
			Col: int32(*col),
		},
	})
	log.Println(resp, err)
}
