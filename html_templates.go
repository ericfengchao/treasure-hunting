package treasure_hunting

import (
	"fmt"
	"strings"
)

const rowTemplate = `
<div class="row">
	%s
</div>
`
const treasureTemplate = `
	<div class="column" style="background-color:#aaa; border-style: solid; border-color: coral;">
    	<p>*</p>
  	</div>
`

const emptySlotTemplate = `
	<div class="column">
    	<p></p>
  	</div>
`
const PlayerTemplate = `
	<div class="column player">
    	<p class="cell">%s</p>
  	</div>
`

const PlayerStatesList = `
<div class="block">
	<div class="row">
		<p>Player: Score </p>
	</div>
	%s
</div>
`

const PlayerStateDiv = `
	<div class="row">
		<div>%s: %d</div>
	</div>
`

const Html = `
<!DOCTYPE html>
<html>
<head>
<title>%s</title>
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
<style>
* {
  font-family: Arial, Helvetica, sans-serif;
  box-sizing: border-box;
}

.left {
  float: left;
  width: 25%%;
}

.right {
  float: right;
  width: 70%%;
}

.treasure {
  background-color: #2164f4;
}

.player {
  background-color: #58ACFA;
}

/* Create two equal columns that floats next to each other */
.column {
  float: left;
  width: 50px;
  height: 50px;
  padding: 1px;
  border-style: solid; 
  border-color: coral; 
  display: flex; 
  align-items: center; 
  justify-content: center;
}

/* Clear floats after the columns */
.row:after {
  content: "";
  display: table;
  clear: both;
  border-color: coral;
}

.cell {
  color: white;
  font-size: 20px;
  
}

</style>
</head>
<body>

<div class="left">
%s
</div>

<div class="right">
%s
</div>
</body>
</html>
`

type BackupViewGameStates struct {
	PlayerStatesView
	Grid *Grid
}

func (s BackupViewGameStates) GetViews() string {
	var allRows []string
	for _, row := range s.Grid.GetSlotRows() {
		var items []string
		for _, slot := range row.GetSlots() {
			slotHtml := emptySlotTemplate
			if slot.GetTreasure() {
				slotHtml = treasureTemplate
			} else if slot.GetPlayerId() != "" {
				slotHtml = fmt.Sprintf(PlayerTemplate, slot.GetPlayerId())
			}
			items = append(items, slotHtml)
		}
		rowHtml := fmt.Sprintf(rowTemplate, strings.Join(items, ""))
		allRows = append(allRows, rowHtml)
	}
	gridView := strings.Join(allRows, "")

	playerStates := s.getPlayerListView()

	return fmt.Sprintf(Html, s.SelfId, playerStates, gridView)
}

type PlayerStatesView struct {
	SelfId       string
	MasterId     string
	SlaveId      string
	PlayerStates []*PlayerState
}

func (p PlayerStatesView) getPlayerListView() string {
	var playerLists []string
	for _, ps := range p.PlayerStates {
		pid := ps.PlayerId
		if pid == p.MasterId {
			pid = "(Primary)" + pid
		} else if pid == p.SlaveId {
			pid = "(Backup)" + pid
		}
		playerM := &PlayerModel{
			id:    pid,
			score: int(ps.Score),
		}
		playerLists = append(playerLists, playerM.getPlayerStateHtml())
	}
	playersHtml := fmt.Sprintf(PlayerStatesList, strings.Join(playerLists, ""))
	return playersHtml
}

type PlayerModeViewGameStates struct {
	gridSize int
	PlayerStatesView
}

func (p PlayerModeViewGameStates) GetViews() string {
	slots := make([][]*slot, p.gridSize)
	for i := range slots {
		slots[i] = make([]*slot, p.gridSize)
		for j := range slots[i] {
			slots[i][j] = new(slot)
		}
	}
	for _, ps := range p.PlayerStates {
		slots[int(ps.CurrentPosition.Row)][int(ps.CurrentPosition.Col)].placePlayer(ps.PlayerId)
	}

	playersHtml := p.getPlayerListView()

	var allRowHtmls []string
	for _, row := range slots {
		var rowItems []string
		for _, s := range row {
			rowItems = append(rowItems, s.getSlotView())
		}
		rowHtml := fmt.Sprintf(rowTemplate, strings.Join(rowItems, ""))
		allRowHtmls = append(allRowHtmls, rowHtml)
	}
	gridView := strings.Join(allRowHtmls, "")

	return fmt.Sprintf(Html, p.SelfId, playersHtml, gridView)

}
