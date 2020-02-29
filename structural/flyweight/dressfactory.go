package main

import (
	"fmt"
	"log"
)

type DressFactory struct {
	dressMap map[string]Dress
}

var DressFactorySingleInstance = &DressFactory{
	dressMap: make(map[string]Dress),
}

const (
	TDress = "TerroristDress"
	CTDress = "CounterTerroristDress"
)

func (factory *DressFactory) GetDress(dressType string) Dress{
	if factory.dressMap[dressType] != nil {
		fmt.Println("found for dressType", dressType)
		return factory.dressMap[dressType]
	}

	fmt.Println("not found for dressType:", dressType)

	switch dressType {
	case TDress:
		factory.dressMap[dressType] = NewTerroristDress("green")
	case CTDress:
		factory.dressMap[dressType] = NewCounterTerroristDress("yellow")
	default:
		log.Fatalf("dressType %v not recognized\n", dressType)
	}

	return factory.dressMap[dressType]
}

func GetDressFactorySingleInstance() *DressFactory {
	return DressFactorySingleInstance
}