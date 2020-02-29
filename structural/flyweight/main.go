package main

import "fmt"

func main() {
	game := &Game{}
	game.AddPlayer(NewPlayer(TDress))
	game.AddPlayer(NewPlayer(TDress))

	game.AddPlayer(NewPlayer(CTDress))
	game.AddPlayer(NewPlayer(CTDress))

	players := game.GetPlayers()

	for _, player := range players {
		fmt.Println("player:", player)
	}
}
