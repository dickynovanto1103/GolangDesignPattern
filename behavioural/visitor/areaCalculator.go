package main

import (
	"fmt"
	"math"
)

type AreaCalculator struct {}

func (a *AreaCalculator) VisitSquare(s Shape) {
	square := s.(*Square)
	area := square.side * square.side
	fmt.Println("area of square: ", area)
}

func (a *AreaCalculator) VisitCircle(s Shape) {
	circle := s.(*Circle)
	area := math.Pi*float64(circle.r*circle.r)
	fmt.Println("area of circle: ", area)
}
