package treasure_hunting

import (
	"fmt"
	"log"
	"strings"
	"sync"
)

type game struct {
	// static states
	grid gridder

	// active states
	rwLock       *sync.RWMutex
	stateVersion int // game state version should be atomically incremented

	// Player management
	playerList map[string]*PlayerModel
}

// based on the new player registry, determine if need to clean up dead players
func (g *game) CleanupPlayer(newPlayerList []*Player) {
	g.rwLock.Lock()
	defer g.rwLock.Unlock()

	newPlayerMap := make(map[string]struct{})
	for _, p := range newPlayerList {
		newPlayerMap[p.GetPlayerId()] = struct{}{}
	}
	for id, p := range g.playerList {
		if _, ok := newPlayerMap[id]; !ok {
			log.Printf("removing player %s", id)
			delete(g.playerList, id)
			g.grid.removePlayer(p.id)
		}
	}
}

func (g *game) GetSerialisedGameStats() *CopyRequest {
	g.rwLock.RLock()
	defer g.rwLock.RUnlock()

	return &CopyRequest{
		Grid:         g.grid.getSerialisedGameStates(),
		PlayerStates: g.GetGameStates(),
		StateVersion: int32(g.stateVersion),
	}
}

// atomic step. Either updated all required information as well as synced with slave Or nothing happened
// atomicity is realised by sync/RWMutex. If game has other member function, Lock/Unlock must be used there as well
func (g *game) MovePlayer(playerId string, move Movement) error {
	g.rwLock.Lock()
	defer g.rwLock.Unlock()

	log.Printf(">>>>>>>.MOVING player %s with MOVEMENT %v", playerId, move)
	// update player
	if p, ok := g.playerList[playerId]; ok {
		// move is received from the endpoint, need listening to the keyboard
		var moveRow, moveCol int
		switch move {
		case Stay:
			return nil
		case West:
			moveRow, moveCol = 0, -1
		case South:
			moveRow, moveCol = 1, 0
		case East:
			moveRow, moveCol = 0, 1
		case North:
			moveRow, moveCol = -1, 0
		}
		newRow := p.currentRow + moveRow
		newCol := p.currentCol + moveCol
		return g.placePlayer(playerId, newRow, newCol)
	} else {
		initialRow, initialCol := g.grid.getRandomEmptySlot()
		return g.placePlayer(playerId, initialRow, initialCol)
	}
}

func (g *game) placePlayer(playerId string, row, col int) error {
	// check if placeable
	if err := g.grid.isPlaceable(row, col); err != nil {
		return err
	}

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
		g.playerList[playerId] = &PlayerModel{
			id:         playerId,
			score:      score,
			currentRow: row,
			currentCol: col,
		}
	}

	// increment version
	g.stateVersion = g.stateVersion + 1

	return nil
}

func (g *game) GetGameStates() []*PlayerState {
	g.rwLock.RLock()
	defer g.rwLock.RUnlock()

	playersSerialised := make([]*PlayerState, 0)
	for _, player := range g.playerList {
		playersSerialised = append(playersSerialised, &PlayerState{
			PlayerId: player.id,
			CurrentPosition: &Coordinate{
				Row: int32(player.currentRow),
				Col: int32(player.currentCol),
			},
			Score: int32(player.score),
		})
	}
	return playersSerialised
}

func (g *game) GetGridView(playerId string) string {
	gridView := g.grid.toGridView()
	playerStates := g.getPlayerStatesListHtml()
	return fmt.Sprintf(Html, playerId, playerStates, gridView)
}

func (g *game) getPlayerStatesListHtml() string {
	var players []string
	for _, p := range g.playerList {
		players = append(players, p.getPlayerStateHtml())
	}
	return fmt.Sprintf(PlayerStatesList, strings.Join(players, ""))
}

func NewGameFromGameCopy(copy *CopyRequest) Gamer {
	fmt.Println("recover from game copy", copy)
	gridCopy := copy.GetGrid()
	originSlot := make([][]*slot, len(gridCopy.GetSlotRows()))
	for i, row := range gridCopy.GetSlotRows() {
		originSlot[i] = make([]*slot, len(row.GetSlots()))
		for j, item := range row.GetSlots() {
			s := &slot{
				treasure: item.Treasure,
				playerId: item.PlayerId,
			}
			originSlot[i][j] = s
		}
	} // convert game_pb.Slot to models.slot

	treasureSlots := make([]int, len(gridCopy.GetTreasureSlots()))
	for i := range gridCopy.GetTreasureSlots() {
		treasureSlots[i] = int(gridCopy.GetTreasureSlots()[i])
	}
	emptySlots := make([]int, len(gridCopy.GetEmptySlots()))
	for i := range gridCopy.GetEmptySlots() {
		emptySlots[i] = int(gridCopy.GetEmptySlots()[i])
	}
	playerSlots := make(map[string]int)
	for playerId, slotId := range gridCopy.GetPlayerSlots() {
		playerSlots[playerId] = int(slotId)
	}
	playerStates := make(map[string]*PlayerModel)
	for _, p := range copy.GetPlayerStates() {
		playerStates[p.PlayerId] = &PlayerModel{
			id:         p.PlayerId,
			score:      int(p.Score),
			currentRow: int(p.CurrentPosition.GetRow()),
			currentCol: int(p.CurrentPosition.GetCol()),
		}
	}

	return &game{
		grid:         NewGridWithGameCopy(originSlot, treasureSlots, playerSlots, emptySlots),
		rwLock:       &sync.RWMutex{},
		stateVersion: int(copy.GetStateVersion()),
		playerList:   playerStates,
	}
}

func NewGame(gridSize int, treasureAmount int) Gamer {
	return &game{
		grid:       newGrid(gridSize, gridSize, treasureAmount),
		rwLock:     &sync.RWMutex{},
		playerList: make(map[string]*PlayerModel),
	}
}
