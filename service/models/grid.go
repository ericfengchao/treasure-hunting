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
	return strings.Join(allRows, "")
}
