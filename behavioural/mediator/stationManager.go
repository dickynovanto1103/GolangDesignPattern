package main

import (
	"fmt"
)

type StationManager struct {
	IsStationFree bool
	TrainQueue []Train
}

func (s *StationManager) CanArrive(train Train) bool {
	if !s.IsStationFree {
		fmt.Println("station is not free now")
		return false
	}
	s.IsStationFree = false
	s.AddTrain(train)
	fmt.Println("train added")
	return true
}

func (s *StationManager) NotifyFree() {
	s.IsStationFree = true
	if len(s.TrainQueue) == 0 {
		fmt.Println("no train queueing")
		return
	}
	s.RemoveTrain()
}

func (s *StationManager) AddTrain(train Train) {
	s.TrainQueue = append(s.TrainQueue, train)
}

func (s *StationManager) RemoveTrain() {
	s.TrainQueue = s.TrainQueue[1:]
}