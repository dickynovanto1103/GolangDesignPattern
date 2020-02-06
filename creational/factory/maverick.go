package main

type maverick struct {
	gunInterface
}

func NewMaverick() *maverick {
	gun := &gun{}
	gun.SetName("maverick")
	gun.SetPower(4)
	return &maverick{gun}
}