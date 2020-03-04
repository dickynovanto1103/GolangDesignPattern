package main

import "GolangDesignPattern/creational/abstractfactory/furniture"

type TraditionalFurnitureFactory struct{}

func (t *TraditionalFurnitureFactory) CreateTable() furniture.Furniture {
	return furniture.NewTable("traditional table")
}

func (t *TraditionalFurnitureFactory) CreateChair() furniture.Furniture {
	return furniture.NewChair("traditional chair")
}