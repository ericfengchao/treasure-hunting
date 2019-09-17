package service

import (
	game_pb "github.com/ericfengchao/treasure-hunting/protos"
	"net/http"
)

type Role string

const (
	PrimaryNode Role = "Primary"
	BackupNode  Role = "Backup"
	PlayerNode  Role = "Player"
)

type GameService interface {
	game_pb.GameServiceServer
	http.Handler
}
