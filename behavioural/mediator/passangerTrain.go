package main

import (
	"fmt"
)

type PassengerTrain struct{
	med Mediator
}

func (t *PassengerTrain) RequestArrival() {
	fmt.Println("passenger train requesting arrival")
	if !t.med.CanArrive(t) {
		fmt.Println("passenger train must wait")
		return
	}
	fmt.Println("passenger train can arrive")
}

func (t *PassengerTrain) Departure() {
	fmt.Println("passenger train want to depart")
	t.med.NotifyFree()
}