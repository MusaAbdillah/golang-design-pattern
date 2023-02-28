package main

import "fmt"

type MiddleCoordinates struct {
	x int
	y int
}

func (a *MiddleCoordinates) visitForSquare(s *Square) {
	fmt.Println("Calculating area for square")
}

func (a *MiddleCoordinates) visitForCircle(c *Circle) {
	fmt.Println("Calculating area for circle")
}

func (a *MiddleCoordinates) visitForRectangle(r *Rectangle) {
	fmt.Println("Calculating area for rectangle")
}
