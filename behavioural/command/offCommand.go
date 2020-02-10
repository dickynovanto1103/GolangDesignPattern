package main

type offCommand struct {
	device device
}

func (off *offCommand) Execute() {
	off.device.Off()
}