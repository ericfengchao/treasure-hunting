package models

import (
	"math/rand"
	"time"
)

type game struct {
	grid           *Grid
	treasureAmount int
}

func (g *game) GetGridView() string {
	return g.grid.ToGridView()
}

func generateNUniqueRandomNumbers(n int, max int) []int {
	res := make([]int, max)
	for i := 0; i < max; i++ {
		res[i] = i
	}
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		r := rand.Intn(max - i)
		res[max-i-1], res[r] = res[r], res[max-i-1]

	}
	return res[:n]
}

func NewGame(gridSize int, treasureAmount int) Gamer {
	grid := &Grid{}
	for i := 0; i < gridSize; i++ {
		grid.slots = append(grid.slots, make([]Slot, gridSize))
	}
	rands := generateNUniqueRandomNumbers(treasureAmount, gridSize*gridSize)
	for _, r := range rands {
		i := r / gridSize
		j := r % gridSize
		grid.slots[i][j].treasure = true
	}
	return &game{
		grid:           grid,
		treasureAmount: treasureAmount,
	}
}

type Gamer interface {
	GetGridView() string
}
