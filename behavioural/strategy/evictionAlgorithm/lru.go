package evictionAlgorithm

import "fmt"

type LRU struct {}

func (algo *LRU) Evict() {
	fmt.Println("eviction using LRU")
}