package main

import "GolangDesignPattern/behavioural/command/command"

type button struct{
	com command.Command
}

func (i *button) Press() {
	i.com.Execute()
}

func (i *button) SetCommand(com command.Command) {
	i.com = com
}