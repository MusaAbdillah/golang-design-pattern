package main

import "fmt"

func main() {
	careTaker := &Caretaker{
		mementoArray: make([]*Memento, 0),
	}

	originator := &Originator{
		state: "A",
	}

	fmt.Printf("Originator current state %s\n", originator.getState())
	careTaker.addMemento(originator.createMemento())

	originator.setState("B")
	fmt.Printf("Originator current state: %s\n", originator.getState())
	careTaker.addMemento(originator.createMemento())

	originator.setState("C")
	fmt.Printf("Originator current state: %s\n", originator.getState())
	careTaker.addMemento(originator.createMemento())

	originator.restoreMemento(careTaker.getMemento(1))
	fmt.Printf("Restored to state: %s\n", originator.getState())

	originator.restoreMemento(careTaker.getMemento(2))
	fmt.Printf("Restore to state: %s\n", originator.getState())
}
