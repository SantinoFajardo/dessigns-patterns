package main

// Car es el producto final que será construido
type Car struct {
	Engine string
	Tires  string
	Seats  int
	Color  string
}

// CarBuilder es la interfaz del constructor
type CarBuilder interface {
	SetEngine(engine string) CarBuilder
	SetTires(tires string) CarBuilder
	SetSeats(seats int) CarBuilder
	SetColor(color string) CarBuilder
	Build() Car
}

// ConcreteCarBuilder es el constructor concreto que implementa la interfaz CarBuilder
type ConcreteCarBuilder struct {
	engine string
	tires  string
	seats  int
	color  string
}
