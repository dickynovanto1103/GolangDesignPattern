package main

type button struct{
	com command
}

func (i *button) Press() {
	i.com.Execute()
}

func (i *button) SetCommand(com command) {
	i.com = com
}