package main

import (
	"GolangDesignPattern/behavioural/bridge/device"
	"GolangDesignPattern/behavioural/bridge/remote"
)

func main() {
	tv := &device.TV{}
	radio := &device.Radio{}

	rem := &remote.Remote{}

	rem.SetDevice(tv)
	rem.PowerOn()
	rem.PowerOff()
	rem.PowerOn()
	rem.IncreaseVolume()
	rem.IncreaseVolume()
	rem.DecreaseVolume()

	rem.SetDevice(radio)
	rem.PowerOn()
	rem.PowerOff()
	rem.PowerOn()
	rem.IncreaseVolume()
	rem.IncreaseVolume()
	rem.IncreaseVolume()
	rem.DecreaseVolume()
}
