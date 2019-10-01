package service

import (
	"net/http"

	game_pb "github.com/ericfengchao/treasure-hunting/protos"
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
