package observer

import "fmt"

type Fans struct {
	ID string
}

func NewFans(ID string) *Fans {
	return &Fans{
		ID: ID,
	}
}

func (f *Fans) Update(id string) {
	fmt.Println("updating fans with id: ", f.ID, "about event with id: ", id)
}

func (f *Fans) GetID() string {
	return f.ID
}
