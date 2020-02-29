package main

func main() {
	game := &Game{}
	game.AddPlayer(NewPlayer(TDress))
	game.AddPlayer(NewPlayer(TDress))

	game.AddPlayer(NewPlayer(CTDress))
	game.AddPlayer(NewPlayer(CTDress))
}
