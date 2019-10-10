package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"

	tracker "github.com/ericfengchao/treasure-hunting/protos"
	"google.golang.org/grpc"
)

var address string = "0.0.0.0"

type server struct {
	PlayerList []*tracker.Player
	RWLock     *sync.RWMutex
	Version    int32
	N          int32
	K          int32
	StartPort  int32
}

func (s *server) Register(ctx context.Context, in *tracker.RegisterRequest) (*tracker.RegisterResponse, error) {
	s.RWLock.Lock()
	defer s.RWLock.Unlock()

	if s.Registered(in.PlayerId) {
		return &tracker.RegisterResponse{
			Status: tracker.RegisterResponse_REGISTERED,
			Registry: &tracker.Registry{
				PlayerList: s.PlayerList,
				Version:    s.Version,
			},
		}, nil
	}

	s.AppendPlayer(&tracker.Player{
		Ip:       address,
		Port:     s.StartPort,
		PlayerId: in.PlayerId,
	})

	res := &tracker.RegisterResponse{
		Status: tracker.RegisterResponse_OK,
		Registry: &tracker.Registry{
			PlayerList: s.PlayerList,
			Version:    s.Version,
		},
		N:            s.N,
		K:            s.K,
		AssignedPort: s.StartPort - 2,
	}
	log.Println("Player: " + in.PlayerId + " registered")
	fmt.Println("TOP player", s.PlayerList[0].PlayerId, s.PlayerList[0].Port)
	fmt.Println("full list", s.PlayerList)
	fmt.Println("============")
	return res, nil

}

func (s *server) AppendPlayer(player *tracker.Player) {
	s.PlayerList = append(s.PlayerList, player)
	s.StartPort += 2 //
	s.Version++
}

func (s *server) ReportMissing(ctx context.Context, in *tracker.Missing) (*tracker.MissingResponse, error) {
	s.RWLock.Lock()
	defer s.RWLock.Unlock()

	if !s.Exist(in.PlayerId) {
		return &tracker.MissingResponse{
			Status: tracker.MissingResponse_NOT_EXIST,
			Registry: &tracker.Registry{
				PlayerList: s.PlayerList,
				Version:    s.Version,
			},
		}, nil
	}

	s.Delete(in.PlayerId)
	log.Println("Delete: " + in.PlayerId)
	if len(s.PlayerList) > 0 {
		fmt.Println("TOP player", s.PlayerList[0].PlayerId, s.PlayerList[0].Port)
		fmt.Println("full list", s.PlayerList)
	} else {
		fmt.Println("NO MORE PLAYERS")
	}
	fmt.Println("============")
	return &tracker.MissingResponse{
		Status: tracker.MissingResponse_OK,
		Registry: &tracker.Registry{
			PlayerList: s.PlayerList,
			Version:    s.Version,
		},
	}, nil
}

func (s *server) Delete(playerId string) {
	for k, v := range s.PlayerList {
		if v.PlayerId == playerId {
			s.PlayerList = append(s.PlayerList[:k], s.PlayerList[k+1:]...)
		}
	}
	s.Version++
} // not include RWLock, but needed

func (s *server) Registered(playerId string) bool {
	for _, v := range s.PlayerList {
		if v.PlayerId == playerId {
			fmt.Println("Player " + playerId + " has registered already")
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
	return &server{
		RWLock:    &sync.RWMutex{},
		N:         n,
		K:         k,
		StartPort: int32(51000),
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
	trackerServer := NewTrackerServer(int32(i32N), int32(i32K))
	svr := grpc.NewServer()
	tracker.RegisterTrackerServiceServer(svr, trackerServer)
	fmt.Println(BANNER)
	svr.Serve(grpcListener)
}
