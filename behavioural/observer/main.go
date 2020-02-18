package main

import (
	"GolangDesignPattern/behavioural/observer/observer"
	"GolangDesignPattern/behavioural/observer/subject"
	"fmt"
)

func main() {
	fans1 := observer.NewFans("fans 1")
	fans2 := observer.NewFans("fans 2")
	fans3 := observer.NewFans("fans 3")

	footballEvent := subject.NewFootballEvent("chelsea vs man utd")
	footballEvent.AddObserver(fans1)
	footballEvent.AddObserver(fans2)
	footballEvent.AddObserver(fans3)

	footballEvent.NotifyObservers()

	fmt.Println("fans2 unsubscribe")
	footballEvent.RemoveObserver(fans2)
	footballEvent.NotifyObservers()

	fmt.Println("fans3 unsubscribe")
	footballEvent.RemoveObserver(fans3)
	footballEvent.NotifyObservers()
}