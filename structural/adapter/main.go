package main

import (
	"fmt"
)

type Mac struct{}

func (m *Mac) InsertSquareUSB() {
	fmt.Println("inserting into square usb")
}

func main() {
	mac := &Mac{}
	fmt.Println(mac)
	mac.InsertSquareUSB()
}
