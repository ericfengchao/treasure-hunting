package models

import (
	"fmt"
	"strings"
)

type Grid struct {
	slots [][]Slot
}

type Slot struct {
	treasure bool   // true means treasure, false means no treasure
	player   string // empty string means non player taken
}

var rowTemplate = `
<div class="row">
	%s
</div>
`
var treasureTemplate = `
	<div class="column" style="background-color:#aaa; border-style: solid; border-color: coral;">
    	<p>Treasure</p>
  	</div>
`

var emptySlotTemplate = `
	<div class="column" style="border-style: solid; border-color: coral;">
    	<p></p>
  	</div>
`
var PlayerTemplate = `
	<div class="column" style="background-color:#ccc; border-style: solid; border-color: coral;">
    	<p>%s</p>
  	</div>
`

const Html = `
<!DOCTYPE html>
<html>
<head>
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
<style>
* {
  box-sizing: border-box;
}

/* Create two equal columns that floats next to each other */
.column {
  float: left;
  width: 100px;
  height: 100px;
  padding: 1px;
}

/* Clear floats after the columns */
.row:after {
  content: "";
  display: table;
  clear: both;
  border-color: coral;
}

</style>
</head>
<body>
	%s
</body>
</html>
`

func (g *Grid) ToGridView() string {
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
	body := strings.Join(allRows, "")
	return fmt.Sprintf(Html, body)
}
