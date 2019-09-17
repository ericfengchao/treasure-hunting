package models

import (
	"fmt"
	"strings"
)

type slot struct {
	treasure bool   // true means treasure, false means no treasure
	player   string // empty string means non player taken
}

func (s *slot) isOccupied() bool {
	return s.player != ""
}

func (s *slot) hasTreasure() bool {
	return s.treasure
}

func (s *slot) placePlayer(playerId string) {
	s.treasure = false // player will take the treasure here if any
	s.player = playerId
}

func (s *slot) removePlayer() {
	s.player = ""
}

func (s *slot) placeTreasure() {
	s.treasure = true
}

type grid struct {
	slots [][]*slot
}

func (g *grid) placeTreasure(treasurePlace int) {
	row := treasurePlace / len(g.slots)
	col := treasurePlace % len(g.slots)
	g.slots[row][col].placeTreasure()
}

func (g *grid) isPlaceable(row, col int) error {
	if row < 0 || row >= len(g.slots) || col < 0 || col >= len(g.slots) {
		return InvalidCoordinates
	}
	s := g.slots[row][col]
	if s.isOccupied() {
		return PlaceAlreadyTaken
	}
	return nil
}

func (g *grid) placePlayer(playerId string, row, col int) bool {
	var huntedTreasure bool
	if g.slots[row][col].hasTreasure() {
		huntedTreasure = true
	}
	g.slots[row][col].placePlayer(playerId)
	return huntedTreasure
}

func (g *grid) removePlayer(row, col int) {
	g.slots[row][col].removePlayer()
}

func (g *grid) toGridView() string {
	var allRows []string
	for _, row := range g.slots {
		var items []string
		for _, slot := range row {
			slotHtml := emptySlotTemplate
			if slot.treasure {
				slotHtml = treasureTemplate
			} else if slot.player != "" {
				slotHtml = fmt.Sprintf(PlayerTemplate, slot.player)
			}
			items = append(items, slotHtml)
		}
		rowHtml := fmt.Sprintf(rowTemplate, strings.Join(items, ""))
		allRows = append(allRows, rowHtml)
	}
	return strings.Join(allRows, "")
}

func newGrid(row, col int, treasures []int) gridder {
	g := &grid{
		slots: make([][]*slot, row),
	}
	for i := 0; i < row; i++ {
		g.slots[i] = make([]*slot, col)
		for j := 0; j < col; j++ {
			g.slots[i][j] = &slot{}
		}
	}
	for _, r := range treasures {
		i := r / col
		j := r % col
		g.slots[i][j].placeTreasure()
	}
	return g
}

type gridder interface {
	toGridView() string
	isPlaceable(row, col int) error

	placePlayer(playerId string, row, col int) bool
	removePlayer(row, col int)
	placeTreasure(treasurePlace int)
}
