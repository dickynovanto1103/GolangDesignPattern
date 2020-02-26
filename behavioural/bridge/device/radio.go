package device

import "fmt"

type Radio struct{
	volume int
}

func (r *Radio) TurnOn() {
	fmt.Println("turning on Radio")
}

func (r *Radio) TurnOff() {
	fmt.Println("turning off Radio")
}

func (r *Radio) DecreaseVolume() {
	fmt.Println("decreasing radio volume")
	r.volume--
	fmt.Println("volume now: ", r.volume)
}

func (r *Radio) IncreaseVolume() {
	fmt.Println("increasing radio volume")
	r.volume++
	fmt.Println("volume now: ", r.volume)
}