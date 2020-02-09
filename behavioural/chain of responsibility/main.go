package main

import "fmt"

func main(){
	patient1 := &patient{
		name: "dicky",
	}
	cashier := &cashier{}
	doctor := &doctor{}
	receptionist := &receptionist{}

	doctor.setNext(cashier)
	receptionist.setNext(doctor)

	fmt.Printf("patient %v entered hospital\n", patient1.name)
	receptionist.execute(patient1)
	receptionist.execute(patient1)

	patient2 := &patient{
		name: "andre",
		hasRegistered: true,
	}
	receptionist.execute(patient2)
}
