package main

type Departement interface {
	execute(*Patient)
	setNext(Departement)
}
