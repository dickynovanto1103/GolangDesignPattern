package main

import "GolangDesignPattern/creational/abstractfactory/furniture"

type FurnitureFactory interface {
	CreateTable() furniture.Furniture
	CreateChair() furniture.Furniture
}

func CreateFactory(factoryType string) FurnitureFactory {
	switch factoryType {
	case "traditional":
		return &TraditionalFurnitureFactory{}
	case "modern":
		return &ModernFurnitureFactory{}
	default:
		panic("factory type: " + factoryType + " is unknown")
	}
}