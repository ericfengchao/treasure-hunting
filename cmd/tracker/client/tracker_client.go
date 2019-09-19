package main

import (
	"context"
	"flag"
	"log"
	"os"

	tracker_pb "github.com/ericfengchao/treasure-hunting/protos/tracker"
	"google.golang.org/grpc"
)

var address string = "localhost:50055"

func main() {
	if len(os.Args) < 3 {
		log.Println("Wrong Args Number")
	}

	ipaddress := os.Args[0]
	port := os.Args[1]
	playerId := os.Args[2]

	flag.Parse()
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed to connect to server", err)
	}
	defer conn.Close()
	client := tracker_pb.NewTrackerServiceClient(conn)
	resp, err := client.Register(context.Background(), &tracker_pb.RegisterRequest{
		Ip:       ipaddress,
		Port:     port,
		PlayerId: playerId,
	})
	// player, _ := strconv.ParseInt(*playerId, 10, 32)
	// resp2, _ := client.ReportMissing(context.Background(), &tracker_pb.Missing{
	// 	PlayerId: int32(player),
	// })
	log.Println(resp, err)
	// log.Println(resp2)
}
