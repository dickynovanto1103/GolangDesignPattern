package main

import "fmt"

func main() {
	caretaker := &caretaker{}
	listString := []string{"a", "ab", "ac"}
	ori := &Originator{}
	for _, val := range listString {
		ori.setState(val)
		fmt.Println("state of originator now:", ori.getState())
		memento := ori.createMemento()
		fmt.Println("crate memento of the originator")
		caretaker.addMemento(memento)
	}

	for idx := range listString {
		memento := caretaker.getMemento(idx)
		fmt.Printf("memento %v has state %v\n", idx, ori.restoreMemento(memento))
	}
}
