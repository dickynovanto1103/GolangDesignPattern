package main

import "fmt"

type MiddleCoordinateCalculator struct {

}

func (m *MiddleCoordinateCalculator) VisitSquare(s Shape) {
	fmt.Println("computing middle coordinate of square")
}

func (m *MiddleCoordinateCalculator) VisitCircle(s Shape) {
	fmt.Println("computing middle coordinate of circle")
}
