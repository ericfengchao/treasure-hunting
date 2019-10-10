package treasure_hunting

import (
	"errors"
	"fmt"
)

type Gamer interface {
	// read
	GetGameStates() []*PlayerState
	GetGridView() string
	GetSerialisedGameStats() *CopyRequest

	// write
	MovePlayer(playerId string, move Movement) error
	CleanupPlayer(playerList []*Player)
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

type PlayerModel struct {
	id         string // 2 char identifier
	score      int
	currentRow int
	currentCol int
}

func (p PlayerModel) getPlayerStateHtml() string {
	return fmt.Sprintf(PlayerStateDiv, p.id, p.score)
}

func (p *PlayerModel) ToPlayerProto() *PlayerState {
	if p == nil {
		return nil
	}
	return &PlayerState{
		PlayerId: p.id,
		Score:    int32(p.score),
		CurrentPosition: &Coordinate{
			Row: int32(p.currentRow),
			Col: int32(p.currentCol),
		},
	}
}
