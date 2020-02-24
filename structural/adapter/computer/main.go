package main

import "GolangDesignPattern/structural/adapter/computer"

func main() {
	mac := &computer.Mac{}

	mac.InsertSquareUSB()

	windowsMachine := &computer.WindowsComputer{}
	windowsAdapter := &computer.WindowsAdapter{
		windowsMachine: windowsMachine,
	}
	windowsAdapter.InsertSquareUSB()
}
