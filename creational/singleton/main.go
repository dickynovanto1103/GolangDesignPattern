package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		go getInstance()
	}
	fmt.Scanln()
}
