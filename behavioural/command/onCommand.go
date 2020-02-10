package main

type onCommand struct{
	device device
}

func (on *onCommand) Execute() {
	on.device.On()
}
