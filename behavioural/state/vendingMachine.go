package main

type VendingMachine struct {
	ItemCount int
	ItemPrice int

	CurrentState State
}

func NewVendingMachine(itemCount, itemPrice int) *VendingMachine {
	vendingMachine := &VendingMachine{
		ItemCount:    itemCount,
		ItemPrice:    itemPrice,
	}

	if itemCount == 0 {
		vendingMachine.CurrentState = NewNoItemState(vendingMachine)
		return vendingMachine
	}

	vendingMachine.CurrentState = NewHasItemState(vendingMachine)
	return vendingMachine
}

func (v *VendingMachine) SetState(state State) {
	v.CurrentState = state
}

func (v *VendingMachine) AddItemCount(count int) {
	v.ItemCount += count
	if v.ItemCount > 0 {
		v.SetState(&HasItemState{})
	}
}