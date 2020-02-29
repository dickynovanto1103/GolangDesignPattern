package main

import dress2 "GolangDesignPattern/structural/flyweight/dress"

type Player struct {
	dress dress2.Dress
}

func NewPlayer(dressType string) *Player {
	dress := dress2.GetDressFactorySingleInstance().GetDress(dressType)
	return &Player{dress: dress}
}
