package state

type State interface {
	RequestItem() error
	AddItem(count int) error
	InsertMoney(money int) error
	DispenseItem() error
}


