package factory

// Abstract factory
type AbstractFurnitureFactory interface {
	// This method return a differnt types of chairs, depending the factory used
	CreateChair() Chair
	// This method return a differnt types of tables, depending the factory used
	CreateTable() Table
}

// Abstract product Chair
type Chair interface {
	SitOn()
}

// Abstract product Table
type Table interface {
	PutFoodOnIt()
}
