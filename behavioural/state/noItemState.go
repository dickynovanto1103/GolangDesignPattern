package main

import (
	"errors"
	"fmt"
)

type NoItemState struct {
	VendingMachine *VendingMachine
}

func NewNoItemState(vendingMachine *VendingMachine) *NoItemState {
	return &NoItemState{VendingMachine: vendingMachine}
}

func (s *NoItemState) RequestItem() error {
	fmt.Println("cannot request item on no item state")
	return errors.New("cannot request item on no item state")
}

func (s *NoItemState) AddItem(count int) error {
	s.VendingMachine.AddItemCount(count)
	return nil
}

func (s *NoItemState) InsertMoney(money int) error {
	fmt.Println("there is no use inserting money here")
	return errors.New("insert money fail in no item state")
}

func (s *NoItemState) DispenseItem() error {
	fmt.Println("cannot dispense item as it is in no item state")
	return errors.New("cannot dispense item in no item state")
}