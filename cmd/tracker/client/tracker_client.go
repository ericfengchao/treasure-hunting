package main

import (
	"context"
	"flag"
	"log"
	"os"

	tracker_pb "github.com/ericfengchao/treasure-hunting/protos/tracker"
	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) < 4 {
		log.Println("Wrong Args Number")
	}

	ipaddress := os.Args[1] // tracker's ip address
	port := os.Args[2]      // tracker's port
	playerId := os.Args[3]  // player's id

	address := ipaddress + ":" + port // concat address

	flag.Parse()
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed to connect to server", err)
	}
	defer conn.Close()
	client := tracker_pb.NewTrackerServiceClient(conn)
	resp, err := client.Register(context.Background(), &tracker_pb.RegisterRequest{
		PlayerId: playerId,
	})
	// resp2, _ := client.ReportMissing(context.Background(), &tracker_pb.Missing{
	// 	PlayerId: playerId,
	// })
	log.Println(resp, err)
	// log.Println(resp2)
}
