package main

import "fmt"

type HPBuilder interface{
	SetMerk()
	SetCountry()
	GetHP() *HP
}

func getHPBuilder(merk string) HPBuilder{
	if merk == "samsung" {
		return &Samsung{}
	} else if merk == "xiaomi" {
		return &Xiaomi{}
	} else {
		fmt.Println("merk not available")
		return nil
	}
}