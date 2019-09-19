package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	tracker "github.com/ericfengchao/treasure-hunting/protos/tracker"
	"google.golang.org/grpc"
)

var address string = "localhost"

type server struct {
	PlayerList []*tracker.Player
	Version    int32
	N          int32
	K          int32
	Port       string
}

func (s *server) Register(ctx context.Context, in *tracker.RegisterRequest) (*tracker.RegisterResponse, error) {
	if s.Registered(in.Ip, in.Port) {
		return &tracker.RegisterResponse{
			Status:  tracker.RegisterResponse_REGISTERED,
			Version: s.Version,
		}, nil
	}

	s.PlayerList = append(s.PlayerList, &tracker.Player{
		Ip:       in.Ip,
		Port:     in.Port,
		PlayerId: in.PlayerId,
	})

	s.Version++
	return &tracker.RegisterResponse{
		Status:     tracker.RegisterResponse_OK,
		PlayerList: s.PlayerList,
		Version:    s.Version,
		N:          s.N,
		K:          s.K,
		Port:       s.Port,
	}, nil
}

func (s *server) ReportMissing(ctx context.Context, in *tracker.Missing) (*tracker.MissingResponse, error) {
	if !s.Exist(in.PlayerId) {
		return &tracker.MissingResponse{
			Status:     tracker.MissingResponse_NOT_EXIST,
			PlayerList: s.PlayerList,
			Version:    s.Version,
		}, nil
	}

	s.Delete(in.PlayerId)
	fmt.Println("Delete: " + in.PlayerId)
	s.Version++

	return &tracker.MissingResponse{
		Status:     tracker.MissingResponse_OK,
		PlayerList: s.PlayerList,
		Version:    s.Version,
	}, nil
}

func (s *server) Delete(playerId string) {
	for k, v := range s.PlayerList {
		if v.PlayerId == playerId {
			s.PlayerList = append(s.PlayerList[:k], s.PlayerList[k+1:]...)
		}
	}
} // not include RWLock, but needed

func (s *server) Registered(Ip, Port string) bool {
	for _, v := range s.PlayerList {
		if v.Ip == Ip && v.Port == Port {
			return true
		}
	}
	return false
}

func (s *server) Exist(playerId string) bool {
	for _, v := range s.PlayerList {
		if v.PlayerId == playerId {
			return true
		}
	}
	return false
}
func NewTrackerServer(n, k int32) *server {
	var playerlist []*tracker.Player
	return &server{
		PlayerList: playerlist,
		Version:    1,
		N:          n,
		K:          k,
		Port:       "50054",
	}
}

func main() {
	if len(os.Args) != 3 {
		log.Println("Wrong param numbers hint:[port][N][K]")
		return
	}

	port := os.Args[0]
	N := os.Args[1]
	K := os.Args[2]

	fullAddress := address + ":" + port
	grpcListener, err := net.Listen("tcp", fullAddress)
	if err != nil {
		log.Fatalf("Failed to listen for grpc: %v", err)
	}
	i32N, _ := strconv.ParseInt(N, 10, 32)
	i32K, _ := strconv.ParseInt(K, 10, 32)
	tracker_server := NewTrackerServer(int32(i32N), int32(i32K))
	svr := grpc.NewServer()
	tracker.RegisterTrackerServiceServer(svr, tracker_server)
	svr.Serve(grpcListener)
}
