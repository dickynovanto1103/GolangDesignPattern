package main

import "fmt"

type doctor struct{
	next department
}

func (d *doctor) execute(p *patient) {
	if !p.hasChecked {
		fmt.Printf("patient %v will be checked\n", p.name)
		p.hasChecked = true
	}
	fmt.Printf("patient %v has been checked by doctor\n", p.name)
	d.next.execute(p)
}

func (d *doctor) setNext(dep department) {
	d.next = dep
}
