package mediator

import (
	"GolangDesignPattern/behavioural/mediator/train"
)

type Mediator interface{
	CanArrive(train.Train) bool
	NotifyFree()
}
