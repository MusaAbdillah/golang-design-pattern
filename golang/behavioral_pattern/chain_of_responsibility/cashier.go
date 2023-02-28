package main

import "fmt"

type Cashier struct {
	next Departement
}

func (c *Cashier) execute(patient *Patient) {
	if patient.paymentDone {
		fmt.Printf("\nPayment already done\n")
	}

	fmt.Printf("\nCashier getting money from patient patient\n")
}

func (c *Cashier) setNext(departement Departement) {
	c.next = departement
}
