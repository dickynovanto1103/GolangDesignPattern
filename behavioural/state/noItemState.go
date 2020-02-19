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
	panic("implement me")
}

func (s *NoItemState) AddItem(count int) error {
	panic("implement me")
}

func (s *NoItemState) InsertMoney(money int) error {
	panic("implement me")
}

func (s *NoItemState) DispenseItem() error {
	fmt.Println("cannot dispense item as it is in no item state")
	return errors.New("cannot dispense item in no item state")
}