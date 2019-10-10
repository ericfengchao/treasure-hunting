package main

import (
	"context"
	"fmt"
	"github.com/ericfengchao/treasure-hunting"
	"log"
	"net"
	"os"
	"strconv"
	"sync"

	"google.golang.org/grpc"
)

var address string = "0.0.0.0"

const BANNER = `
.___________..______          ___       ______  __  ___  _______ .______
|           ||   _  \        /   \     /      ||  |/  / |   ____||   _  \
\---|  |----\|  |_)  |      /  ^  \   |  ,----'|  '  /  |  |__   |  |_)  |
    |  |     |      /      /  /_\  \  |  |     |    <   |   __|  |      /
    |  |     |  |\  \----./  _____  \ |  \----.|  .  \  |  |____ |  |\  \----.
    |__|     | _| \._____/__/     \__\ \______||__|\__\ |_______|| _| \._____|
`

type server struct {
	PlayerList []*treasure_hunting.Player
	RWLock     *sync.RWMutex
	Version    int32
	N          int32
	K          int32
	StartPort  int32
}

func (s *server) Register(ctx context.Context, in *treasure_hunting.RegisterRequest) (*treasure_hunting.RegisterResponse, error) {
	s.RWLock.Lock()
	defer s.RWLock.Unlock()

	if s.Registered(in.PlayerId) {
		return &treasure_hunting.RegisterResponse{
			Status: treasure_hunting.RegisterResponse_REGISTERED,
			Registry: &treasure_hunting.Registry{
				PlayerList: s.PlayerList,
				Version:    s.Version,
			},
		}, nil
	}

	s.AppendPlayer(&treasure_hunting.Player{
		Ip:       address,
		Port:     s.StartPort,
		PlayerId: in.PlayerId,
	})

	res := &treasure_hunting.RegisterResponse{
		Status: treasure_hunting.RegisterResponse_OK,
		Registry: &treasure_hunting.Registry{
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

func (s *server) AppendPlayer(player *treasure_hunting.Player) {
	s.PlayerList = append(s.PlayerList, player)
	s.StartPort += 2 //
	s.Version++
}

func (s *server) ReportMissing(ctx context.Context, in *treasure_hunting.Missing) (*treasure_hunting.MissingResponse, error) {
	s.RWLock.Lock()
	defer s.RWLock.Unlock()

	if !s.Exist(in.PlayerId) {
		return &treasure_hunting.MissingResponse{
			Status: treasure_hunting.MissingResponse_NOT_EXIST,
			Registry: &treasure_hunting.Registry{
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
	return &treasure_hunting.MissingResponse{
		Status: treasure_hunting.MissingResponse_OK,
		Registry: &treasure_hunting.Registry{
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
	treasure_hunting.RegisterTrackerServiceServer(svr, trackerServer)
	fmt.Println(BANNER)
	svr.Serve(grpcListener)
}
