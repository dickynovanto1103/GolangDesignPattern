package main

import "fmt"

type HPBuilder interface{
	Reset()
	SetMerk()
	SetCountry()
}

func getHPBuilder(merk string) HPBuilder{
	if merk == "samsung" {
		samsungBuilder := &SamsungBuilder{}
		samsungBuilder.Reset()
		return samsungBuilder
	} else if merk == "xiaomi" {
		xiaomiBuilder := &XiaomiBuilder{}
		xiaomiBuilder.Reset()
		return xiaomiBuilder
	} else {
		fmt.Println("merk not available")
		return nil
	}
}