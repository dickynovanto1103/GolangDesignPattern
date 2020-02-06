package main

import "fmt"

func main() {
	gunFactory := &gunFactory{}
	ak47, err := gunFactory.createGun("ak47")
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("ak47, attr:", ak47.GetName(), ak47.GetPower())

	maverick, err := gunFactory.createGun("maverick")
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("maverick, attr:", maverick.GetName(), maverick.GetPower())
}
