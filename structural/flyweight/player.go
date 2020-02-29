package main

type Player struct {
	dress      Dress
}

func NewPlayer(dressType string) *Player {
	dress := GetDressFactorySingleInstance().GetDress(dressType)
	return &Player{dress: dress}
}
