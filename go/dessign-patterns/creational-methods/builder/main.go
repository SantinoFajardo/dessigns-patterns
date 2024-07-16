package main

import "fmt"

func main() {
	builder := NewCarBuilder()
	director := NewDirector(builder)

	sportCar := director.BuildSportCar()
	familyCar := director.BuildFamilyCar()

	fmt.Printf("Auto Deportivo: %+v\n", sportCar)
	fmt.Printf("Auto Familiar: %+v\n", familyCar)
}
