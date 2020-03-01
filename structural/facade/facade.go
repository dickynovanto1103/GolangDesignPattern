package main

import (
	"GolangDesignPattern/structural/facade/iteminfo"
	"GolangDesignPattern/structural/facade/itemstock"
	"errors"
	"fmt"
)

type OrderSystemFacade struct {}

func (s *OrderSystemFacade) OrderItem(item string) error {
	fmt.Println("order item: ", item)
	itemInfo := &iteminfo.ItemInfo{}
	exist := itemInfo.CheckItemInfo(item)
	if !exist {
		return errors.New("item " + item + " does not exist in item info DB")
	}

	itemStock := &itemstock.ItemStock{}
	stock := itemStock.CheckItemStock(item)
	if stock == 0 {
		return errors.New("item " + item + "does not have stock left")
	}
	return nil
}