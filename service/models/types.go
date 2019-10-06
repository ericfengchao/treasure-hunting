package models

import (
	"errors"

	game_pb "github.com/ericfengchao/treasure-hunting/protos"
)

type Gamer interface {
	// read
	GetGameStates() []*game_pb.PlayerState
	GetGridView() string
	GetSerialisedGameStats() *game_pb.CopyRequest

	// write
	MovePlayer(playerId string, move Movement) error
	CleanupPlayer(playerList []*game_pb.Player)
}

var (
	InvalidCoordinates = errors.New("invalid coordinates")
	PlaceAlreadyTaken  = errors.New("place is already taken")
	SlaveIsDown        = errors.New("slave is down")
	GridIsFull         = errors.New("grid is full")
)

type Role string

const (
	PrimaryNode Role = "Primary"
	BackupNode  Role = "Backup"
	PlayerNode  Role = "Player"
)

type Movement int

const (
	Stay Movement = iota
	West
	South
	East
	North
)

const Exit Movement = 9
