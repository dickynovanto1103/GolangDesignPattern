package main

import "fmt"

type cashier struct{}

func (c *cashier) execute(p *patient) {
	if !p.hasPaid {
		fmt.Printf("patient %v want to pay\n", p.name)
		p.hasPaid = true
	}
	fmt.Printf("patient %v has paid\n", p.name)

}

func (c *cashier) setNext(d department) {
}