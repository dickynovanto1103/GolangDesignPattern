package main

type gunInterface interface {
	SetName(name string)
	SetPower(power int)
	GetName() string
	GetPower() int
}
