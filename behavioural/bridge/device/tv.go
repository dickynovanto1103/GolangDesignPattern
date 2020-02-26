package device

import "fmt"

type TV struct{
	volume int
}

func (t *TV) TurnOn() {
	fmt.Println("turning on TV")
}

func (t *TV) TurnOff() {
	fmt.Println("turning off TV")
}

func (t *TV) DecreaseVolume() {
	fmt.Println("decreasing volume of TV")
	t.volume--
	fmt.Println("volume now: ", t.volume)
}

func (t *TV) IncreaseVolume() {
	fmt.Println("increasing volume of TV")
	t.volume++
	fmt.Println("volume now: ", t.volume)
}
