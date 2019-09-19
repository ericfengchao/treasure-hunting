package main

import (
	"context"
	"flag"
	"log"

	tracker_pb "github.com/ericfengchao/treasure-hunting/protos/tracker"
	"google.golang.org/grpc"
)

var address string = "localhost:50055"

var ipaddress = flag.String("ip", "127.0.0.1", "Pass your ip address to tracker")
var port = flag.String("port", "23", "Pass your port number to tracker")
var playerId = flag.String("id", "1", "Who is missing")

func main() {
	flag.Parse()
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed to connect to server", err)
	}
	defer conn.Close()
	client := tracker_pb.NewTrackerServiceClient(conn)
	resp, err := client.Register(context.Background(), &tracker_pb.RegisterRequest{
		Ip:   *ipaddress,
		Port: *port,
	})
	// player, _ := strconv.ParseInt(*playerId, 10, 32)
	// resp2, _ := client.ReportMissing(context.Background(), &tracker_pb.Missing{
	// 	PlayerId: int32(player),
	// })
	log.Println(resp, err)
	// log.Println(resp2)
}
