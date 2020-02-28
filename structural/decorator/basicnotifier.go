package main

import "fmt"

type BasicNotifier struct {}

func (n *BasicNotifier) Notify() {
	fmt.Println("basic notifier Notify")
}