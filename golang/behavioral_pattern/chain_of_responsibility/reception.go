package main

import "fmt"

type Reception struct {
	next Departement
}

func (r *Reception) execute(patient *Patient) {
	if patient.registrationDone {
		fmt.Printf("\nPatient registration already done\n")
		r.next.execute(patient)
		return
	}

	fmt.Printf("\nReception registering patient\n")
	patient.registrationDone = true
	r.next.execute(patient)
}

func (r *Reception) setNext(departement Departement) {
	r.next = departement
}
