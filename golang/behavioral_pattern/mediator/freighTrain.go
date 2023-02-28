package main

import "fmt"

type FreighTrain struct {
	mediator Mediator
}

func (f *FreighTrain) arrive() {
	if !f.mediator.canArrive(f) {
		fmt.Println("FreighTrain: Arrival blocked, waiting!")
		return
	}

	fmt.Println("FreighTrain: Arrived")
}

func (f *FreighTrain) depart() {
	fmt.Println("FreighTrain: Leaving")
	f.mediator.notifyAboutDeparture()
}

func (f *FreighTrain) permitArrival() {
	fmt.Println("FreighTrain: Arrival permitted")
	f.arrive()
}
