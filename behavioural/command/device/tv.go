package device

import "fmt"

type Tv struct{}

func (t *Tv) On() {
	fmt.Println("tv on")
}

func (t *Tv) Off() {
	fmt.Println("tv off")
}
