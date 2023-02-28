package main

import "fmt"

type Tv struct {
	isRunning bool
}

func (t *Tv) on() {
	t.isRunning = true
	fmt.Printf("\nTurning tv on\n")
}

func (t *Tv) off() {
	t.isRunning = false
	fmt.Printf("\nTurning tv off\n")
}
