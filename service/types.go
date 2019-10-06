package service

import (
	"net/http"

	game_pb "github.com/ericfengchao/treasure-hunting/protos"
)

type GameService interface {
	game_pb.GameServiceServer
	http.Handler
	GetLocalRegistry() *game_pb.Registry
	UpdateLocalRegistry(*game_pb.Registry)
}
