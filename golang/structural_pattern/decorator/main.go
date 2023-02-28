package main

import "fmt"

func main() {
	pizza := &VeggeMania{}

	//add cheese topping
	pizzaWithCheese := &CheeseToping{
		pizza: pizza,
	}

	//add tomatto toping
	pizzaWithCheeseAndTomatto := &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggemania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomatto.getPrice())
}
