package main

func main() {
	mac := &Mac{}

	mac.InsertSquareUSB()

	windowsMachine := &WindowsComputer{}
	windowsAdapter := &WindowsAdapter{
		windowsMachine: windowsMachine,
	}
	windowsAdapter.InsertSquareUSB()
}
