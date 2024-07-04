package factory

// Interface definiction of the animal. This interface will be used by all the animals created
type Animal interface {
	// Method of the Animal interface that all animals implementation implement.
	Speak() string
}

// Dog is a Animal implementation that implement all the methods required by that interface
type Dog struct{}

// Use the `Speak()` method declared by the 'Animal' interface
func (d *Dog) Speak() string {
	return "Woof!"
}

// Cat is a Animal interface implementation
type Cat struct{}

// Use the `Spek()` method declared by the 'Animal' interface
func (d *Cat) Speak() string {
	return "Woof!"
}
