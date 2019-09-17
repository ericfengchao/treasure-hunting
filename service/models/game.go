package models

import (
	"fmt"
	"strings"
	"sync"
)

type game struct {
	// static states
	grid gridder

	// treasure management
	treasurePlaces []int
	treasureAmount int

	// active states
	rwLock       *sync.RWMutex
	stateVersion int // game state version should be atomically incremented

	// player management
	playerList map[string]*player
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

	// place new treasure if required
	if huntedTreasure {
		//var newTreasure int // TODO generate new treasure coordinates
		//g.grid.placeTreasure(newTreasure)
	}

	// update player states
	if p, ok := g.playerList[playerId]; ok {
		if huntedTreasure {
			p.score = p.score + 1
		}
		// remove player from current pos
		g.grid.removePlayer(p.currentRow, p.currentCol)
		// update pos to latest
		p.currentRow = row
		p.currentCol = col
	} else {
		var score int
		if huntedTreasure {
			score = 1
		}
		g.playerList[playerId] = &player{
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

func NewGame(gridSize int, treasureAmount int) Gamer {
	treasurePlaces := generateNUniqueRandomNumbers(treasureAmount, gridSize*gridSize)
	return &game{
		grid:           newGrid(gridSize, gridSize, treasurePlaces),
		rwLock:         &sync.RWMutex{},
		treasureAmount: treasureAmount,
		treasurePlaces: treasurePlaces,
		playerList:     make(map[string]*player),
	}
}
