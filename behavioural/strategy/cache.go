package main

import (
	"GolangDesignPattern/behavioural/strategy/evictionAlgorithm"
	"fmt"
)

type Cache struct {
	evictAlgorithm evictionAlgorithm.EvictionAlgorithm

	capacity int
	length int
}

func NewCache(length, capacity int) *Cache {
	return &Cache{
		capacity:       capacity,
		length:         length,
	}
}

func (c *Cache) SetEvictAlgorithm(evictionAlgorithm evictionAlgorithm.EvictionAlgorithm) {
	c.evictAlgorithm = evictionAlgorithm
}

func (c *Cache) evict() {
	if c.length == 0 {
		fmt.Println("cannot evict since length == 0")
		return
	}
	c.length--
	c.evictAlgorithm.Evict()
	fmt.Println("success evict")
}

func (c *Cache) add() {
	if c.length == c.capacity {
		c.evict()
	}
	c.length++
	fmt.Println("success add")
}