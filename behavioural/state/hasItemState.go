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
	panic("implement me")
}

func (s *HasItemState) InsertMoney(money int) error {
	panic("implement me")
}

func (s *HasItemState) DispenseItem() error {
	panic("implement me")
}