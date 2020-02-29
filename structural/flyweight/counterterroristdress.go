package main

type CounterTerroristDress struct {
	color string
}

func NewCounterTerroristDress(color string) *CounterTerroristDress {
	return &CounterTerroristDress{color: color}
}

func (d *CounterTerroristDress) GetColor() string {
	return d.color
}