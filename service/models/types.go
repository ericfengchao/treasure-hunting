package models

import "errors"

type Gamer interface {
	// read
	GetGameStates() map[string]*Player
	GetGridView() string

	// write
	PlacePlayer(playerId string, row, col int) (bool, error)
}

var (
	InvalidCoordinates = errors.New("invalid coordinates")
	PlaceAlreadyTaken  = errors.New("place is already taken")
	SlaveIsDown        = errors.New("slave is down")
)
