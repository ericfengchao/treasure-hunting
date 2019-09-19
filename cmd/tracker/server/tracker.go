package main

import (
	"context"
	"fmt"
	"log"
	"net"

	tracker "github.com/ericfengchao/treasure-hunting/protos/tracker"
	"google.golang.org/grpc"
)

type server struct {
	PlayerList []*tracker.Player
	Version    int32
	Count      int32
}

func (s *server) Register(ctx context.Context, in *tracker.RegisterRequest) (*tracker.RegisterResponse, error) {
	if s.Registered(in.Ip, in.Port) {
		return &tracker.RegisterResponse{
			Status:  tracker.RegisterResponse_REGISTERED,
			Version: s.Version,
		}, nil
	}

	playerId := s.Count // for every version's update, this version id can also refered to playerid
	s.Count++
	s.PlayerList = append(s.PlayerList, &tracker.Player{
		Ip:       in.Ip,
		Port:     in.Port,
		PlayerId: playerId,
	})
	s.Version++
	return &tracker.RegisterResponse{
		Status:     tracker.RegisterResponse_OK,
		PlayerList: s.PlayerList,
		Version:    s.Version,
	}, nil
}

func (s *server) ReportMissing(ctx context.Context, in *tracker.Missing) (*tracker.MissingResponse, error) {
	if s.Exist(in.PlayerId) {
		return &tracker.MissingResponse{
			Status:     tracker.MissingResponse_NOT_EXIST,
			PlayerList: s.PlayerList,
			Version:    s.Version,
		}, nil
	} // if s.Exist == false means that player does not exist

	s.Delete(in.PlayerId)
	fmt.Println("Delete: " + string(in.PlayerId))
	s.Version++

	return &tracker.MissingResponse{
		Status:     tracker.MissingResponse_OK,
		PlayerList: s.PlayerList,
		Version:    s.Version,
	}, nil
}

func (s *server) Delete(playerId int32) {
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

func (s *server) Exist(playerId int32) bool {
	for _, v := range s.PlayerList {
		if v.PlayerId == playerId {
			return false
		}
	}
	return true
}
func NewTrackerServer() *server {
	var playerlist []*tracker.Player
	return &server{
		PlayerList: playerlist,
		Version:    1,
		Count:      1,
	}
}

func main() {

	grpcListener, err := net.Listen("tcp", "localhost:50055")
	if err != nil {
		log.Fatalf("Failed to listen for grpc: %v", err)
	}
	tracker_server := NewTrackerServer()
	svr := grpc.NewServer()
	tracker.RegisterTrackerServiceServer(svr, tracker_server)
	svr.Serve(grpcListener)
}
