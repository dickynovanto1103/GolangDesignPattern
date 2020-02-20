package main

import "GolangDesignPattern/behavioural/strategy/evictionAlgorithm"

func main() {
	cache := NewCache(3, 4)
	cache.SetEvictAlgorithm(&evictionAlgorithm.LRU{})
	cache.add()
	cache.add()
	cache.add()
	cache.add()
	cache.SetEvictAlgorithm(&evictionAlgorithm.FIFO{})
	cache.evict()
	cache.evict()
	cache.add()
	cache.add()
	cache.add()
	cache.add()
}
