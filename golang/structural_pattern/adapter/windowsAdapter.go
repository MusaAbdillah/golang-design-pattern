package main

import "fmt"

type WindowsAdapter struct {
	windowMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort(){
	fmt.Println("Adapter converts lightning signal to usb")
	w.windowMachine.insertIntoUSBPort()
}


