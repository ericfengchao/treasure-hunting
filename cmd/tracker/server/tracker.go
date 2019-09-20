package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"

	tracker "github.com/ericfengchao/treasure-hunting/protos/tracker"
	"google.golang.org/grpc"
)

var address string = "localhost"

type server struct {
	PlayerList []*tracker.Player
	RWLock     *sync.RWMutex
	Version    int32
	N          int32
	K          int32
	StartPort  int32
}

func (s *server) Register(ctx context.Context, in *tracker.RegisterRequest) (*tracker.RegisterResponse, error) {

	if s.Registered(in.PlayerId) {
		return &tracker.RegisterResponse{
			Status:  tracker.RegisterResponse_REGISTERED,
			Version: s.Version,
		}, nil
	}

	res, err := s.AppendPlayer(&tracker.Player{
		Ip:       "localhost",
		Port:     s.StartPort,
		PlayerId: in.PlayerId,
	})

	if err != nil {
		log.Println(err)
	}
	return res, nil
}
func (s *server) AppendPlayer(player *tracker.Player) (*tracker.RegisterResponse, error) {
	s.RWLock.Lock()
	defer s.RWLock.Unlock()

	s.PlayerList = append(s.PlayerList, player)

	res := &tracker.RegisterResponse{
		Status:     tracker.RegisterResponse_OK,
		PlayerList: s.PlayerList,
		Version:    s.Version,
		N:          s.N,
		K:          s.K,
		StartPort:  s.StartPort,
	}
	s.StartPort++
	s.Version++
	fmt.Println("Player: " + player.PlayerId + " registered")
	return res, nil
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

	return &tracker.MissingResponse{
		Status:     tracker.MissingResponse_OK,
		PlayerList: s.PlayerList,
		Version:    s.Version,
	}, nil
}

func (s *server) Delete(playerId string) {
	s.RWLock.Lock()
	defer s.RWLock.Unlock()
	for k, v := range s.PlayerList {
		if v.PlayerId == playerId {
			s.PlayerList = append(s.PlayerList[:k], s.PlayerList[k+1:]...)
		}
	}
	s.Version++
} // not include RWLock, but needed

func (s *server) Registered(playerId string) bool {
	s.RWLock.RLock()
	defer s.RWLock.RUnlock()
	for _, v := range s.PlayerList {
		if v.PlayerId == playerId {
			fmt.Println("Player " + playerId + " has registered already")
			return true
		}
	}
	return false
}

func (s *server) Exist(playerId string) bool {
	s.RWLock.RLock()
	defer s.RWLock.RUnlock()
	for _, v := range s.PlayerList {
		if v.PlayerId == playerId {
			return true
		}
	}
	return false
}
func NewTrackerServer(n, k int32) *server {
	var playerlist []*tracker.Player
	var Lock sync.RWMutex
	return &server{
		PlayerList: playerlist,
		RWLock:     &Lock,
		Version:    1,
		N:          n,
		K:          k,
		StartPort:  int32(51000),
	}
}

func main() {
	if len(os.Args) != 4 {
		log.Println("Wrong param numbers hint:[port][N][K]")
		return
	}

	port := os.Args[1]
	N := os.Args[2]
	K := os.Args[3]

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
