package train

import (
	"GolangDesignPattern/behavioural/mediator/mediator"
	"fmt"
)

type GoodsTrain struct{
	med mediator.Mediator
}

func (t *GoodsTrain) RequestArrival() {
	fmt.Println("goods train request to arrive")

	if !t.med.CanArrive(t) {
		fmt.Println("goods train must wait")
	}

	fmt.Println("goods train can enter station")
}

func (t *GoodsTrain) Departure() {
	fmt.Println("good train depart")
	t.med.NotifyFree()
}