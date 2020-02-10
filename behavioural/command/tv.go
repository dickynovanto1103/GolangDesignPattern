package main

import "fmt"

type tv struct{}

func (t *tv) On() {
	fmt.Println("tv on")
}

func (t *tv) Off() {
	fmt.Println("tv off")
}
