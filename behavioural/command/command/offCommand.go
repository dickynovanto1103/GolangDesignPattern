package command

import (
	"GolangDesignPattern/behavioural/command/device"
)

type OffCommand struct {
	Device device.Device
}

func (off *OffCommand) Execute() {
	off.Device.Off()
}