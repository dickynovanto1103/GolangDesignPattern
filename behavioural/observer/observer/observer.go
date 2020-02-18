package observer

type Observer interface {
	Update(id string)
	GetID() string
}
