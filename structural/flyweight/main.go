package main

import (
	"GolangDesignPattern/structural/flyweight/dress"
	"fmt"
)

func main() {
	game := &Game{}
	game.AddPlayer(NewPlayer(dress.TDress))
	game.AddPlayer(NewPlayer(dress.TDress))

	game.AddPlayer(NewPlayer(dress.CTDress))
	game.AddPlayer(NewPlayer(dress.CTDress))

	players := game.GetPlayers()

	for _, player := range players {
		fmt.Println("player:", player)
	}
}
