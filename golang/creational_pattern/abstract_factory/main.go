package main

import "fmt"

func main() {
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	adidasShoe := adidasFactory.makeShoe()
	nikeShoe := nikeFactory.makeShoe()

	adidasShirt := adidasFactory.makeShirt()
	nikeShirt := adidasFactory.makeShirt()

	printShoeDetails(adidasShoe)
	printShoeDetails(nikeShoe)

	printShirtDetails(adidasShirt)
	printShirtDetails(nikeShirt)
}

func printShoeDetails(s IShoe) {
	fmt.Printf("Shoes Logo %v", s.getLogo())
	fmt.Println()

	fmt.Printf("Shoes Size %v", s.getSize())
	fmt.Println()
}

func printShirtDetails(s IShirt) {
	fmt.Printf("Shirt Logo %v", s.getLogo())
	fmt.Println()

	fmt.Printf("Shirt Size %v", s.getSize())
	fmt.Println()
}
