package main

import (
	"log"
	"context"
	"net"
	"google.golang.org/grpc"
	tracker "github.com/ericfengchao/treasure-hunting/tracker/protos"
	"fmt"
)

var playerList []string

type server struct{}

func (s *server) GetInfo(ctx context.Context, in *tracker.InfoRequest) (*tracker.InfoResponse, error) {
	playerList = append(playerList, in.Ip)
	return &tracker.InfoResponse{Players: playerList}, nil
}

func main(){

	grpcListener, err := net.Listen("tcp", "localhost:50055")
	if err != nil {
		log.Fatalf("Failed to listen for grpc: %v", err)
	}

	svr := grpc.NewServer()
	tracker.RegisterTrackerServiceServer(svr, &server{})

	svr.Serve(grpcListener)

	fmt.Println("TRACKER STARTS NOW")
}