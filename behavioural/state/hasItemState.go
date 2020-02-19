package main

import (
	"errors"
	"fmt"
)

type HasItemState struct {
	VendingMachine *VendingMachine
}

func NewHasItemState(vendingMachine *VendingMachine) *HasItemState {
	return &HasItemState{VendingMachine: vendingMachine}
}

func (s *HasItemState) RequestItem() error {
	if s.VendingMachine.ItemCount == 0 {
		fmt.Println("item count == 0")
		return errors.New("item count == 0")
	}

	fmt.Println("requesting item")
	return nil
}

func (s *HasItemState) AddItem(count int) error {
	s.VendingMachine.AddItemCount(count)
	return nil
}

//TODO: keep state of money stored
func (s *HasItemState) InsertMoney(money int) error {
	if s.VendingMachine.ItemPrice > money {
		fmt.Printf("money is not enough, money: %v item Price: %v\n", money, s.VendingMachine.ItemPrice)
		return errors.New("money is not enough, returning back money")
	}
	fmt.Println("success inserting money")
	return nil
}

func (s *HasItemState) DispenseItem() error {
	s.VendingMachine.ItemCount--
	if s.VendingMachine.ItemCount == 0 {
		s.VendingMachine.SetState(NewNoItemState(s.VendingMachine))
	}
	return nil
}