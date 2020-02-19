package main

func main() {
	vendingMachine := NewVendingMachine(1, 1)
	vendingMachine.CurrentState.RequestItem()
	vendingMachine.CurrentState.DispenseItem()
	vendingMachine.CurrentState.RequestItem()
}
