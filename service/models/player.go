package models

import (
	"fmt"
	game_pb "github.com/ericfengchao/treasure-hunting/protos"
)

type Player struct {
	id         string // 2 char identifier
	host       string
	port       int
	score      int
	currentRow int
	currentCol int
}

func (p Player) getPlayerStateHtml() string {
	return fmt.Sprintf(PlayerState, p.id, p.score)
}

func (p *Player) ToPlayerProto() *game_pb.PlayerState {
	if p == nil {
		return nil
	}
	return &game_pb.PlayerState{
		PlayerId: p.id,
		Score:    int32(p.score),
		CurrentPosition: &game_pb.Coordinate{
			Row: int32(p.currentRow),
			Col: int32(p.currentCol),
		},
	}
}
