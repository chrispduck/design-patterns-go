package main

import "fmt"

// Component interface
type IPizza interface {
	getPrice() int
}

//  Base concrete
type TomatoBase struct{}

func (TomatoBase) getPrice() int {
	return 10
}

// Decorator A
type CheeseTopping struct {
	pizza IPizza
}

func (p CheeseTopping) getPrice() int {
	return p.pizza.getPrice() + 2
}

// Decorator B
type PepperoniTopping struct {
	pizza IPizza
}

func (p PepperoniTopping) getPrice() int {
	return p.pizza.getPrice() + 3
}

func main() {
	var pizza IPizza
	pizza = &TomatoBase{}
	pizza = &CheeseTopping{pizza: pizza}
	pizza = &PepperoniTopping{pizza: pizza}
	fmt.Println(pizza.getPrice())
}
