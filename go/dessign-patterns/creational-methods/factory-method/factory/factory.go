package factory

// Factory definition and factory method
type AnimalFactory struct{}

// Method of the factory that always return the Animal interface. Doing so, the client can always manage the
// object retuned by the factory in the same way and doesn't need to have concatenated ifs to handle different
// animals
func (f AnimalFactory) CreateAnimal(t int) Animal {
	switch t {
	case 1:
		return &Dog{}
	case 2:
		return &Cat{}
	default:
		return nil
	}
}
