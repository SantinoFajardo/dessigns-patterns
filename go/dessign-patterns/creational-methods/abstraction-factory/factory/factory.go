package factory

import "fmt"

// ========================
// ==== Modern Factory ====
// =========================

// Concrete Product Modern Chair
type ModernChair struct{}

func NewModernChair() Chair {
	return &ModernChair{}
}

func (c *ModernChair) SitOn() {
	fmt.Println("Sitting on a modern chair...")
}

// Concrete Product Modern Table
type ModernTable struct{}

func (t *ModernTable) PutFoodOnIt() {
	fmt.Println("Putting food on a modern table...")
}

// Concrete Factory for Modern Furniture
// This struct will create objects based on the Modern furnitures
type ModernFurnitureFactory struct{}

func (f *ModernFurnitureFactory) CreateChair() Chair {
	return &ModernChair{}
}

func (f *ModernFurnitureFactory) CreateTable() Table {
	return &ModernTable{}
}

// ===========================
// ==== Victorian Factory ====
// ===========================

// Concrete Product Victorian Chair
type VictorianChair struct{}

func (c *VictorianChair) SitOn() {
	fmt.Println("Sitting on a Victorian chair.")
}

// Concrete Product Victorian Table
type VictorianTable struct{}

func (t *VictorianTable) PutFoodOnIt() {
	fmt.Println("Putting food on a Victorian table.")
}

// Concrete Factory for Victorian Furniture
// This struct will create objects based on the Victorian furnitures
type VictorianFurnitureFactory struct{}

// Create Victorian chairs
func (f *VictorianFurnitureFactory) CreateChair() Chair {
	return &VictorianChair{}
}

// Create Victorian table
func (f *VictorianFurnitureFactory) CreateTable() Table {
	return &VictorianTable{}
}
