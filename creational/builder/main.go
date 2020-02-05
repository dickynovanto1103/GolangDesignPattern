package main

import "fmt"

func main() {
	samsungBuilder := getHPBuilder("samsung")
	director := NewDirector(samsungBuilder)
	samsungHP := director.Build()

	fmt.Println("samsung:", samsungHP.GetHP().merk, samsungHP.GetHP().country)

	xiaomiBuilder := getHPBuilder("xiaomi")
	director = NewDirector(xiaomiBuilder)
	xiaomiHP := director.Build()
	fmt.Println("xiaomi:", xiaomiHP.GetHP().merk, xiaomiHP.GetHP().country)
}
