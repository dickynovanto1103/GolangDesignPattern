package train

import (
	"GolangDesignPattern/behavioural/mediator/mediator"
	"fmt"
)

type PassengerTrain struct{
	med mediator.Mediator
}

func (t *PassengerTrain) RequestArrival() {
	fmt.Println("requesting arrival")
}

func (t *PassengerTrain) Departure() {
	fmt.Println("train want to depart")
}