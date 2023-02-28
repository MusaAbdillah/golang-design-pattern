package main

import "fmt"

func main() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Println("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Println("Normal house Window Type: %s\n", normalHouse.windowType)
	fmt.Println("Normal House Num Floor: %d\n", normalHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Println("Igloo House Door Type: %s\n", iglooHouse.doorType)
	fmt.Println("Igloo House Window type: %s\n", iglooHouse.windowType)
	fmt.Println("Igloo House Num Floor %d\n", iglooHouse.floor)
}
