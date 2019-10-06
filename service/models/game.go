package models

import (
	"fmt"
	"strings"
	"sync"

	game_pb "github.com/ericfengchao/treasure-hunting/protos"
)

type game struct {
	role Role

	// static states
	grid gridder

	// active states
	rwLock       *sync.RWMutex
	stateVersion int // game state version should be atomically incremented

	// Player management
	playerList map[string]*Player
}

// atomic step. Either updated all required information as well as synced with slave Or nothing happened
// atomicity is realised by sync/RWMutex. If game has other member function, Lock/Unlock must be used there as well
func (g *game) PlacePlayer(playerId string, row, col int) (bool, error) {
	g.rwLock.Lock()
	defer g.rwLock.Unlock()

	// check if placeable
	if err := g.grid.isPlaceable(row, col); err != nil {
		return false, err
	}

	// sync with slave
	//var syncedWithSlave bool
	// TODO sync with slave
	//if !syncedWithSlave {
	//	return false, SlaveIsDown
	//}

	// once slave is committed, start writing in master only
	huntedTreasure := g.grid.placePlayer(playerId, row, col)

	// update Player states
	if p, ok := g.playerList[playerId]; ok {
		if huntedTreasure {
			p.score = p.score + 1
		}
		p.currentRow = row
		p.currentCol = col
	} else {
		var score int
		if huntedTreasure {
			score = 1
		}
		g.playerList[playerId] = &Player{
			id:         playerId,
			score:      score,
			currentRow: row,
			currentCol: col,
		}
	}

	// increment version
	g.stateVersion = g.stateVersion + 1

	return huntedTreasure, nil
}

func (g *game) MovePlayer(playerId string, move string) (bool, error) {
	g.rwLock.Lock()
	defer g.rwLock.Unlock()
	// move is received from the endpoint, need listening to the keyboard
	var moveRow, moveCol int
	if move == "0\n" {
		return true, nil
	}
	if move == "1\n" {
		moveRow, moveCol = -1, 0
	}
	if move == "2\n" {
		moveRow, moveCol = 0, 1
	}
	if move == "3\n" {
		moveRow, moveCol = 1, 0
	}
	if move == "4\n" {
		moveRow, moveCol = 0, -1
	}
	// update player

	if p, ok := g.playerList[playerId]; ok {
		newCol := p.currentCol + moveCol
		newRow := p.currentRow + moveRow
		rowsize, colsize := g.grid.getSize()
		// check if placeable
		if newCol > colsize-1 || newCol < 0 {
			return false, InvalidCoordinates
		}
		if newRow > rowsize-1 || newRow < 0 {
			return false, InvalidCoordinates
		} // judge boundary
		if err := g.grid.isPlaceable(newRow, newCol); err != nil {
			return false, err
		}
		huntedTreasure := g.grid.placePlayer(playerId, newCol, newRow)
		if huntedTreasure {
			p.score = p.score + 1
		}
		p.currentRow = newRow
		p.currentCol = newCol
		return huntedTreasure, nil
	} else {
		return false, NoPlayerFound
	}
}

func (g *game) GetGameStates() map[string]*Player {
	g.rwLock.RLock()
	defer g.rwLock.RUnlock()

	return g.playerList
}

func (g *game) GetGridView() string {
	gridView := g.grid.toGridView()
	playerStates := g.getPlayerStatesListHtml()
	return fmt.Sprintf(Html, playerStates, gridView)
}

func (g *game) getPlayerStatesListHtml() string {
	var players []string
	for _, p := range g.playerList {
		players = append(players, p.getPlayerStateHtml())
	}
	return fmt.Sprintf(PlayerStatesList, strings.Join(players, ""))
}

func (g *game) UpdateFullCopy(slots [][]*game_pb.Slot, treasureSlots []int, playerSlots map[string]int, emptySlots []int, stateVersion int) {
	g.rwLock.Lock()
	defer g.rwLock.Unlock()
	g.stateVersion = stateVersion
	originSlot := make([][]*slot, len(slots))
	for i, row := range slots {
		originSlot[i] = make([]*slot, len(row))
		for j, item := range row {
			s := &slot{
				treasure: item.Treasure,
				playerId: item.PlayerId,
			}
			originSlot[i][j] = s
		}
	} // convert game_pb.Slot to models.slot
	g.grid.updateGrid(originSlot, treasureSlots, playerSlots, emptySlots)
}

func NewGame(gridSize int, treasureAmount int, role Role) Gamer {
	return &game{
		role:       role,
		grid:       newGrid(gridSize, gridSize, treasureAmount),
		rwLock:     &sync.RWMutex{},
		playerList: make(map[string]*Player),
	}
}
