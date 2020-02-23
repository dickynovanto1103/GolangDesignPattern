package main

func main() {
	circle := &Circle{r: 10}
	square := &Square{side: 4}
	areaCalc := &AreaCalculator{}
	circle.Accept(areaCalc)
	square.Accept(areaCalc)

	middleCoordCalc := &MiddleCoordinateCalculator{}
	circle.Accept(middleCoordCalc)
	square.Accept(middleCoordCalc)
}
