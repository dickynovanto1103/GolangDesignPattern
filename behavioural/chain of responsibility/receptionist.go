package main

import "fmt"

type receptionist struct {
	next department
}

func (r *receptionist) execute(p *patient) {

	if !p.hasRegistered {
		fmt.Printf("patient %v will register\n", p.name)
		p.hasRegistered = true
	}

	fmt.Printf("patient %v has registered\n", p.name)
	r.next.execute(p)
}

func (r *receptionist) setNext(d department) {
	r.next = d
}