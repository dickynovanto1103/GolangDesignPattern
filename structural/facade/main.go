package main

import (
	"log"
)

func main() {
	orderFacade := &OrderSystemFacade{}
	err := orderFacade.OrderItem("book")
	if err != nil {
		log.Println("err: ", err)
	}

	log.Println("order item book succeed")

	err = orderFacade.OrderItem("printer")
	if err != nil {
		log.Println("err: ", err)
	}
}
