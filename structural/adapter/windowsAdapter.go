package main

type WindowsAdapter struct{
	windowsMachine *WindowsComputer
}

func (w *WindowsAdapter) InsertSquareUSB() {
	w.windowsMachine.InsertCircleUSB()
}

