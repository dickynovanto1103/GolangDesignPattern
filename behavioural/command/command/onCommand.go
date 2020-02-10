package command

import (
	"GolangDesignPattern/behavioural/command/device"
)

type OnCommand struct{
	Device device.Device
}

func (on *OnCommand) Execute() {
	on.Device.On()
}
