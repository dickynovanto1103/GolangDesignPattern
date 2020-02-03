package main

import (
	"fmt"
)

type Mac struct{}

func (m *Mac) InsertSquareUSB() {
	fmt.Println("inserting into square usb")
}

type WindowsAdapter struct{
	windowsMachine *WindowsComputer
}

func (w *WindowsAdapter) InsertSquareUSB() {
	w.windowsMachine.InsertCircleUSB()
}

type WindowsComputer struct{}

func (w *WindowsComputer) InsertCircleUSB() {
	fmt.Println("inserting into circle usb")
}

func main() {
	mac := &Mac{}

	mac.InsertSquareUSB()

	windowsMachine := &WindowsComputer{}
	windowsAdapter := &WindowsAdapter{
		windowsMachine: windowsMachine,
	}
	windowsAdapter.InsertSquareUSB()
}
