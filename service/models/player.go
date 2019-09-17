package models

import "fmt"

type player struct {
	id         string // 2 char identifier
	host       string
	port       int
	score      int
	currentRow int
	currentCol int
}

func (p player) getPlayerStateHtml() string {
	return fmt.Sprintf(PlayerState, p.id, p.score)
}
