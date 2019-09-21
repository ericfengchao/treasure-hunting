package service

import (
	game_pb "github.com/ericfengchao/treasure-hunting/protos"
	"net/http"
)

type GameService interface {
	game_pb.GameServiceServer
	http.Handler
	GetLocalRegistry() *game_pb.Registry
	UpdateLocalRegistry(*game_pb.Registry)
}
