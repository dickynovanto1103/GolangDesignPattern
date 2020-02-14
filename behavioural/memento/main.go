package main

import "fmt"

func main() {
	caretaker := &caretaker{}
	listString := []string{"a", "ab", "ac"}
	for _, val := range listString {
		ori := &Original{}
		ori.setState(val)
		memento := ori.createMemento()
		caretaker.addMemento(memento)
	}

	for idx := range listString {
		memento := caretaker.getMemento(idx)
		fmt.Printf("memento %v has state %v\n", idx, memento.getSavedState())
	}
}
