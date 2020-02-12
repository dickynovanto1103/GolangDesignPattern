package main

import (
	"GolangDesignPattern/behavioural/mediator/mediator"
	"fmt"
)

func main() {
	med := &mediator.StationManager{}
	fmt.Println("med:", med)
}
