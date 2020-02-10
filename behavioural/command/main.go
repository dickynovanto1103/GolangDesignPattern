package main

import (
	"GolangDesignPattern/behavioural/command/command"
	"GolangDesignPattern/behavioural/command/device"
	"fmt"
)

func main() {
	tv := &device.Tv{}
	onCommand := &command.OnCommand{Device: tv}
	offCommand := &command.OffCommand{Device: tv}

	onButton := &button{com: onCommand}
	offButton := &button{com: offCommand}

	fmt.Println("turning on tv")
	onButton.Press()

	fmt.Println("turning off tv")
	offButton.Press()
}
