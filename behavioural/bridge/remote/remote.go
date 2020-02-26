package remote

import "GolangDesignPattern/behavioural/bridge/device"

type Remote struct {
	dev device.Device
}

func (r *Remote) PowerOn() {
	r.dev.TurnOn()
}

func (r *Remote) PowerOff() {
	r.dev.TurnOff()
}

func (r *Remote) DecreaseVolume() {
	r.dev.DecreaseVolume()
}

func (r *Remote) IncreaseVolume() {
	r.dev.IncreaseVolume()
}

func (r *Remote) SetDevice(dev device.Device) {
	r.dev = dev
}