package main

import "fmt"

type Topping struct {
	Name string
}

type Pizza struct {
	Diameter int
	Toppings []*Topping
}

func (p *Pizza) eat() {
	fmt.Println("nom nom nom")
}

func main() {
	pizza := &Pizza{
		Diameter: 12,
		Toppings: make([]*Topping, 0),
	}

	pizza.Toppings = append(pizza.Toppings, &Topping{"Pineapple"})

	fmt.Println(pizza)
	fmt.Println(*pizza)

	fmt.Println(pizza.Toppings[0])
	fmt.Println(*pizza.Toppings[0])

	fmt.Println(pizza.Toppings[0].Name)

	pizza.eat()
}
