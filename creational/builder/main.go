package main

import "fmt"

func main() {
	samsungBuilder := getHPBuilder("samsung")
	director := NewDirector(samsungBuilder)
	director.BuildHP()

	samsungHP := samsungBuilder.(*SamsungBuilder).GetHP()
	fmt.Println("samsung:", samsungHP.merk, samsungHP.country)

	xiaomiBuilder := getHPBuilder("xiaomi")
	director = NewDirector(xiaomiBuilder)
	director.BuildHP()

	xiaomiHP := xiaomiBuilder.(*XiaomiBuilder).GetHP()

	fmt.Println("xiaomi:", xiaomiHP.merk, xiaomiHP.country)
}
