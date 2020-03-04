package main

import "GolangDesignPattern/creational/abstractfactory/furniture"

type ModernFurnitureFactory struct{}

func (m *ModernFurnitureFactory) CreateTable() furniture.Furniture {
	return furniture.NewTable("modern table")
}

func (m *ModernFurnitureFactory) CreateChair() furniture.Furniture {
	return furniture.NewChair("modern chair")
}

