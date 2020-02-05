package main

import "fmt"

func main() {
	samsungBuilder := getHPBuilder("samsung")
	director := NewDirector(samsungBuilder)
	samsungHP := director.BuildHP()

	fmt.Println("samsung:", samsungHP.merk, samsungHP.country)

	xiaomiBuilder := getHPBuilder("xiaomi")
	director = NewDirector(xiaomiBuilder)
	xiaomiHP := director.BuildHP()
	fmt.Println("xiaomi:", xiaomiHP.merk, xiaomiHP.country)
}
