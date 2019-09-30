package main

import (
	"fmt"
	"github.com/ericfengchao/treasure-hunting/player_service"
	"log"
	"os"
)

var gridSize = 3
var treasureAmount = 5

func main() {
	if len(os.Args) < 4 {
		log.Fatal("insufficient params to start the game. [tracker Host] [tracker Port] [Player Id]")
	}
	trackerHost := os.Args[1]
	trackerPort := os.Args[2]
	playerId := os.Args[3]

	playerSvc := player_service.NewPlayerSvc(trackerHost, trackerPort, playerId)
	defer playerSvc.Close()

	go playerSvc.StartServing()

	closing := make(chan struct{}, 0)
	go playerSvc.Start(closing)

	<-closing
	fmt.Printf("PLAYER %s EXITING\n", playerId)
}
