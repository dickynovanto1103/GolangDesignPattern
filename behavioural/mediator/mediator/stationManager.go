package mediator

import (
	"GolangDesignPattern/behavioural/mediator/train"
	"fmt"
)

type StationManager struct {
	IsStationFree bool
	TrainQueue []train.Train
}

func (s *StationManager) CanArrive(train train.Train) bool {
	if !s.IsStationFree {
		fmt.Println("station is not free now")
		return false
	}
	s.IsStationFree = false
	AddTrain(train)
	fmt.Println("train added")
	return true
}

func (s *StationManager) NotifyFree() {
	s.IsStationFree = true
	if len(s.TrainQueue) == 0 {
		fmt.Println("no train queueing")
		return
	}
	RemoveTrain()
}

func AddTrain(train train.Train) {
	s.TrainQueue = append(s.TrainQueue, train)
}

func RemoveTrain() {
	s.TrainQueue = s.TrainQueue[1:]
}