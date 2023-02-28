package main

import "fmt"

type Doctor struct {
	next Departement
}

func (d *Doctor) execute(patient *Patient) {
	if patient.doctorCheckUpDone {
		fmt.Printf("\nDoctor checkup already done\n")
		d.next.execute(patient)
		return
	}

	fmt.Printf("\nDoctor checking up patient\n")
	patient.doctorCheckUpDone = true
	d.next.execute(patient)
}

func (d *Doctor) setNext(departement Departement) {
	d.next = departement
}
