package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		go getInstance()
	}
	fmt.Scanln()

	for i := 0; i < 100; i++ {
		go getInsOnceInstance()
	}
	fmt.Scanln()
}
