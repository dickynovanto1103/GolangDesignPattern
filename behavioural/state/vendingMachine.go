package state

type VendingMachine struct {
	ItemCount int
	ItemPrice int

	CurrentState State
}

func newVendingMachine(itemCount, itemPrice int) *VendingMachine {
	vendingMachine := &VendingMachine{
		ItemCount:    itemCount,
		ItemPrice:    itemPrice,
	}

	if itemCount == 0 {
		vendingMachine.CurrentState = &NoItemState{}
		return vendingMachine
	}

	vendingMachine.CurrentState = &HasItemState{}
	return vendingMachine
}