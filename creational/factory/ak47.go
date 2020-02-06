package main

type ak47 struct {
	gunInterface
}

func NewAK47() *ak47 {
	gun := &gun{}
	gun.SetName("ak47")
	gun.SetPower(1)
	return &ak47{gun}
}