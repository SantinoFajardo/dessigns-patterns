package main

import "github.com/santinofajardo/abstract-factory/factory"

func main() {
	// Factory used by the client to create furnitures. It doesn't matters if is a modern
	// factory or victorian, the client will handle all as the same since follow the same
	// interface
	var furnitureFactory factory.AbstractFurnitureFactory

	// Create Modern Furniture
	// This factory will just create modern furnitures.
	// But the client doesn't matters if is modern or not
	furnitureFactory = &factory.ModernFurnitureFactory{}
	chair := furnitureFactory.CreateChair()
	table := furnitureFactory.CreateTable()

	// Use the modern furnitures
	chair.SitOn()
	table.PutFoodOnIt()

	// Create Victorian Furniture
	// This factory will just create victorian furnitures
	// But the client doesn't matters if is modern or not
	// Take in consideration that instance is just called once
	furnitureFactory = &factory.VictorianFurnitureFactory{}
	chair = furnitureFactory.CreateChair()
	table = furnitureFactory.CreateTable()

	// Use Victorian furnitures
	chair.SitOn()
	table.PutFoodOnIt()
}
