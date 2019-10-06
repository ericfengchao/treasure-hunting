package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ericfengchao/treasure-hunting/player_service"
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
	//KeyboardListen
	// playerSvc.Initialize()
	go playerSvc.StartServing()

	closing := make(chan struct{}, 0)
	//go playerSvc.Start(closing)
	go playerSvc.KeyboardListen(closing)
	<-closing
	fmt.Printf("PLAYER %s EXITING\n", playerId)
}
