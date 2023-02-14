package main

import "fmt"

// ---- ABSTRACTION ----

// Interface
type VehicleIface interface {
	Run() string
	Start() string
}
func StartVehicle(v VehicleIface) string {
	return v.Start()
}
func RunVehicle(v VehicleIface) string {
	return v.Run()
}

// Super type
type Vehicle struct {
	Name  string
	Brand string
	Year  int
}

// ---- INHERITANCE / POLYMORPHISM ----

// Struct
type Car struct {
	Vehicle
	// Car specific fields
	Doors int
	HasFourWheels bool
}
type Motorcycle struct {
	Vehicle
	// Motorcycle specific fields
	HasTwoWheels bool
	Handlebar string
}

// Methods --> automatic implementation of the interface
func (c Car) Run() string {
	return "The car is running..."
}
func (c Car) Start() string {
	return "Starting " + c.Name + "..."
}

func (m Motorcycle) Run() string {
	return "The motorcycle is running..."
}
func (m Motorcycle) Start() string {
	return "Starting " + m.Name + "..."
}

func main() {
	myCar := Car{
		Vehicle: Vehicle{
			Name:  "Gol",
			Brand: "Volkswagen",
			Year:  2010,
		},
		Doors: 4,
		HasFourWheels: true,
	}

	myMotorcycle := Motorcycle{
		Vehicle: Vehicle{
			Name:  "Ninja",
			Brand: "Kawasaki",
			Year:  2010,
		},
		HasTwoWheels: true,
		Handlebar: "Sport",
	}

	fmt.Println(myCar)
	fmt.Println(StartVehicle(myCar))
	fmt.Println(RunVehicle(myCar))
	fmt.Println()
	fmt.Println(myMotorcycle)
	fmt.Println(StartVehicle(myMotorcycle))
	fmt.Println(RunVehicle(myMotorcycle))
}