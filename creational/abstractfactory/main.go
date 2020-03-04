package main

import "fmt"

func main() {
	traditionalFactory := CreateFactory("traditional")
	modernFactory := CreateFactory("modern")

	client := NewClient(traditionalFactory)
	chair := client.furnitureFactory.CreateChair()
	table := client.furnitureFactory.CreateTable()

	fmt.Println("chair: " + chair.GetName())
	fmt.Println("table:" + table.GetName())

	client.SetFurnitureFactory(modernFactory)
	chair = client.furnitureFactory.CreateChair()
	table = client.furnitureFactory.CreateTable()

	fmt.Println("chair: " + chair.GetName())
	fmt.Println("table:" + table.GetName())
}