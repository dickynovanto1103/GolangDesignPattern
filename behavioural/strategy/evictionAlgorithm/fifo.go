package evictionAlgorithm

import "fmt"

type FIFO struct {}

func (algo *FIFO) Evict() {
	fmt.Println("eviction using FIFO")
}
