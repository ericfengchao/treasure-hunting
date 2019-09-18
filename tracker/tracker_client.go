package main

import (
	tracker_pb "github.com/ericfengchao/treasure-hunting/tracker/protos"
	"flag"
	"google.golang.org/grpc"
	"log"
	"context"
)

var address string = "localhost:50055"

var ipaddress = flag.String("ip","127.0.0.1","Pass your ip address to tracker")

func main() {
	flag.Parse()
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed to connect to server", err)
	}
	defer conn.Close()
	client := tracker_pb.NewTrackerServiceClient(conn)
	resp, err := client.GetInfo(context.Background(), &tracker_pb.InfoRequest{Ip: *ipaddress})
	log.Println(resp, err)
}
