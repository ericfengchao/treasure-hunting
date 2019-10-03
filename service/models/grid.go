package models

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

type slot struct {
	treasure bool   // true means treasure, false means no treasure
	playerId string // empty string means non Player taken
}

func (s *slot) isOccupied() bool {
	return s.playerId != ""
}

func (s *slot) hasTreasure() bool {
	return s.treasure
}

func (s *slot) placePlayer(playerId string) {
	s.treasure = false // Player will take the treasure here if any
	s.playerId = playerId
}

func (s *slot) removePlayer() {
	s.playerId = ""
}

func (s *slot) placeTreasure() {
	s.treasure = true
}

type grid struct {
	slots [][]*slot

	// indices for fast retrieval
	treasureSlots []int // positions of treasure hiding slots
	playerSlots   map[string]int
	emptySlots    []int
}

func (g *grid) getRowCol(pos int) (row int, col int) {
	numOfColumns := len(g.slots[0])
	return pos / numOfColumns, pos % numOfColumns
}

func (g *grid) getPos(row, col int) int {
	numOfColumns := len(g.slots[0])
	return row*numOfColumns + col
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

	newTreasurePos := -1
	// update indices
	newPos := g.getPos(row, col)
	if origPos, exists := g.playerSlots[playerId]; exists {
		g.playerSlots[playerId] = newPos
		origRow, origCol := g.getRowCol(origPos)
		g.slots[origRow][origCol].removePlayer()

		if huntedTreasure {
			candidates := append(g.emptySlots, origPos)
			newTreasurePos = candidates[rand.Intn(len(candidates))]
			g.emptySlots = removeIntFromSlice(candidates, newTreasurePos)
			g.treasureSlots = replaceXWithY(g.treasureSlots, newPos, newTreasurePos)
		} else {
			g.emptySlots = replaceXWithY(g.emptySlots, newPos, origPos)
		}

	} else {
		g.playerSlots[playerId] = newPos

		if huntedTreasure {
			newTreasurePos = g.emptySlots[rand.Intn(len(g.emptySlots))]
			g.emptySlots = removeIntFromSlice(g.emptySlots, newTreasurePos)
			g.treasureSlots = replaceXWithY(g.treasureSlots, newPos, newTreasurePos)
		} else {
			g.emptySlots = removeIntFromSlice(g.emptySlots, newPos)
		}
	}
	if newTreasurePos != -1 {
		newTreasureRow, newTreasureCol := g.getRowCol(newTreasurePos)
		g.slots[newTreasureRow][newTreasureCol].placeTreasure()
	}

	log.Println("============DEBUG==========")
	log.Println(g.treasureSlots)
	log.Println(g.playerSlots)
	log.Println(g.emptySlots)
	log.Println("============DEBUG==========")

	return huntedTreasure
}

func removeIntFromSlice(orig []int, k int) []int {
	var new []int
	for _, num := range orig {
		if num != k {
			new = append(new, num)
		}
	}
	return new
}

func replaceXWithY(orig []int, x, y int) []int {
	var new []int
	for _, num := range orig {
		if num == x {
			new = append(new, y)
		} else {
			new = append(new, num)
		}
	}
	return new
}

func (g *grid) removePlayer(playerId string) {
	if origPos, exists := g.playerSlots[playerId]; exists {
		row, col := g.getRowCol(origPos)
		g.slots[row][col].removePlayer()
		g.emptySlots = append(g.emptySlots, origPos)
	}
}

func (g *grid) toGridView() string {
	var allRows []string
	for _, row := range g.slots {
		var items []string
		for _, slot := range row {
			slotHtml := emptySlotTemplate
			if slot.treasure {
				slotHtml = treasureTemplate
			} else if slot.playerId != "" {
				slotHtml = fmt.Sprintf(PlayerTemplate, slot.playerId)
			}
			items = append(items, slotHtml)
		}
		rowHtml := fmt.Sprintf(rowTemplate, strings.Join(items, ""))
		allRows = append(allRows, rowHtml)
	}
	return strings.Join(allRows, "")
}

func (g *grid) updateGrid(slots [][]*slot, treasureSlots []int, playerSlots map[string]int, emptySlots []int) {
	g.slots = slots
	g.emptySlots = emptySlots
	g.playerSlots = playerSlots
	g.treasureSlots = treasureSlots
}

func (g *grid) getSize() (int, int) {
	return len(g.slots), len(g.slots[0])
}

func newGrid(row, col int, treasureAmount int) gridder {
	rand.Seed(time.Now().Unix())
	shuffledN := rand.Perm(row * col)
	g := &grid{
		slots:         make([][]*slot, row),
		treasureSlots: shuffledN[:treasureAmount],
		playerSlots:   make(map[string]int),
		emptySlots:    shuffledN[treasureAmount:],
	}
	for i := 0; i < row; i++ {
		g.slots[i] = make([]*slot, col)
		for j := 0; j < col; j++ {
			g.slots[i][j] = &slot{}
		}
	}
	for _, r := range g.treasureSlots {
		i := r / col
		j := r % col
		g.slots[i][j].placeTreasure()
	}
	return g
}

type gridder interface {
	toGridView() string
	isPlaceable(row, col int) error
	updateGrid(slots [][]*slot, treasureSlots []int, playerSlots map[string]int, emptySlots []int)
	placePlayer(playerId string, row, col int) bool
	removePlayer(playerId string)
	getSize() (int, int)
}
