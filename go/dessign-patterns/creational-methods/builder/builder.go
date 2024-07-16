package main

// NewCarBuilder crea una nueva instancia de ConcreteCarBuilder
func NewCarBuilder() *ConcreteCarBuilder {
	return &ConcreteCarBuilder{}
}

func (b *ConcreteCarBuilder) SetEngine(engine string) CarBuilder {
	b.engine = engine
	return b
}

func (b *ConcreteCarBuilder) SetTires(tires string) CarBuilder {
	b.tires = tires
	return b
}

func (b *ConcreteCarBuilder) SetSeats(seats int) CarBuilder {
	b.seats = seats
	return b
}

func (b *ConcreteCarBuilder) SetColor(color string) CarBuilder {
	b.color = color
	return b
}

func (b *ConcreteCarBuilder) Build() Car {
	return Car{
		Engine: b.engine,
		Tires:  b.tires,
		Seats:  b.seats,
		Color:  b.color,
	}
}

// Director define el orden de construcción del objeto
type Director struct {
	builder CarBuilder
}

// NewDirector crea una nueva instancia de Director
func NewDirector(builder CarBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

// BuildSportCar construye un auto deportivo
func (d *Director) BuildSportCar() Car {
	return d.builder.
		SetEngine("V8").
		SetTires("Deportivas").
		SetSeats(2).
		SetColor("Rojo").
		Build()
}

// BuildFamilyCar construye un auto familiar
func (d *Director) BuildFamilyCar() Car {
	return d.builder.
		SetEngine("V6").
		SetTires("All-Season").
		SetSeats(5).
		SetColor("Azul").
		Build()
}
