package main

import "fmt"

type Medical struct {
	next Departement
}

func (m *Medical) execute(patient *Patient) {
	if patient.medicineDone {
		fmt.Printf("\nPatient medicine already done\n")
		m.next.execute(patient)
		return
	}

	fmt.Printf("\nMedical giving medicine to patient\n")
	m.next.execute(patient)
}

func (m *Medical) setNext(departement Departement) {
	m.next = departement
}
