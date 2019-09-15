package main

import (
	"flag"
	"fmt"
	"github.com/ericfengchao/treasure-hunting/service/models"
	"log"
)

var gridSize = flag.Int("gridSize", 3, "size of the Square Grid")
var treasureAmount = flag.Int("treasure", 5, "number of treasures in grid")

func main() {
	flag.Parse()
	if *treasureAmount > (*gridSize)*(*gridSize) {
		log.Println("treasure amount is greater than grid. Automatically set it to be the size of the grid")
		*treasureAmount = *gridSize
	}
	log.Printf("GridSize: %d, TreasureAmount: %d\n", *gridSize, *treasureAmount)
	game := models.NewGame(*gridSize, *treasureAmount)
	fmt.Println(game.GetGridView())
}
