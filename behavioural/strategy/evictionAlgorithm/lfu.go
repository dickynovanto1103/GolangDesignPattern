package evictionAlgorithm

import "fmt"

type LFU struct {}

func (algo *LFU) Evict() {
	fmt.Println("eviction using LFU")
}
