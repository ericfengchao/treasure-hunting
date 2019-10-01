package models

import (
	"errors"

	game_pb "github.com/ericfengchao/treasure-hunting/protos"
)

type Gamer interface {
	// read
	GetGameStates() map[string]*Player
	GetGridView() string
	// write
	PlacePlayer(playerId string, row, col int) (bool, error)
	UpdateFullCopy(slots [][]*game_pb.Slot, treasureSlots []int, playerSlots map[string]int, emptySlots []int, stateVersion int)
}

var (
	InvalidCoordinates = errors.New("invalid coordinates")
	PlaceAlreadyTaken  = errors.New("place is already taken")
	SlaveIsDown        = errors.New("slave is down")
)
