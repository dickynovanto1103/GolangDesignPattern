package main

import "fmt"

func main() {
	tv := &tv{}
	onCommand := &onCommand{device: tv}
	offCommand := &offCommand{device: tv}

	onButton := &button{com: onCommand}
	offButton := &button{com: offCommand}

	fmt.Println("turning on tv")
	onButton.Press()

	fmt.Println("turning off tv")
	offButton.Press()
}
