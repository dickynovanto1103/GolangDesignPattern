package main

type Game struct {
	players []*Player
}

func (g *Game) AddPlayer(p *Player) {
	g.players = append(g.players, p)
}

func (g *Game) GetPlayers() []*Player {
	return g.players
}