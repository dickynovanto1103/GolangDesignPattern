package main

import (
	"fmt"
)

func main() {
	med := &StationManager{}
	med.NotifyFree()
	fmt.Println("med:", med)
	pt := &PassengerTrain{med: med}
	gt := &GoodsTrain{med: med}
	pt.RequestArrival()
	gt.RequestArrival()

	fmt.Println()
	pt.Departure()
	gt.RequestArrival()

	pt1 := &PassengerTrain{med: med}
	pt2 := &PassengerTrain{med: med}
	fmt.Println()
	pt1.RequestArrival()
	pt2.RequestArrival()

	fmt.Println()
	gt.Departure()
	pt1.RequestArrival()
	pt2.RequestArrival()

	fmt.Println()
	pt1.Departure()
	pt2.RequestArrival()
}
