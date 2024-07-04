package main

import (
	"fmt"

	"github.com/santinofajardo/factory-method/factory"
)

// Client code
func main() {
	factory := factory.AnimalFactory{}

	dog := factory.CreateAnimal(1)
	fmt.Println("Dog talk: ", dog.Speak())

	cat := factory.CreateAnimal(2)
	fmt.Println("Cat tal: ", cat.Speak())
}
