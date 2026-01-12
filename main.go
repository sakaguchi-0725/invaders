package main

import (
	"invaders/game"
	"log"
)

func main() {
	game, err := game.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}
