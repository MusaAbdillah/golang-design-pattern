package main

import "fmt"

type VeggeMania struct {
}

func (v *VeggeMania) getPrice() int {
	fmt.Printf("Initialize price was 15 \n")
	return 15
}
