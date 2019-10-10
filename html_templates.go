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
    	<p>⭐️</p>
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

type ViewableGameStats struct {
	Grid         *Grid
	PlayerStates []*PlayerState
}

func (s ViewableGameStats) GetGridView() string {
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

	var players []string
	for _, p := range s.PlayerStates {
		player := &PlayerModel{
			id:    p.PlayerId,
			score: int(p.Score),
		}
		players = append(players, player.getPlayerStateHtml())
	}
	playerStates := fmt.Sprintf(PlayerStatesList, strings.Join(players, ""))

	return fmt.Sprintf(Html, playerStates, gridView)
}
