package main

import "fmt"

type CheeseToping struct {
	pizza IPizza
}

func (c *CheeseToping) getPrice() int {
	pizzaPrice := c.pizza.getPrice() + 10
	fmt.Printf("WithTomatoTopping price was %d \n", pizzaPrice)
	return pizzaPrice
}
