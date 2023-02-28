package main

import "fmt"

type TomatoTopping struct {
	pizza IPizza
}

func (t *TomatoTopping) getPrice() int {
	pizzaPrice := t.pizza.getPrice() + 7
	fmt.Printf("WithTomatoTopping price was %d \n", pizzaPrice)
	return pizzaPrice
}
