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

// ---- INHERITANCE ----

// Struct
type Car struct {
	Vehicle
}

// Methods --> automatic implementation of the interface
func (c Car) Run() string {
	return "Running..."
}
func (c Car) Start() string {
	return "Starting " + c.Name + "..."
}

func main() {
	myCar := Car{
		Vehicle: Vehicle{
			Name:  "Gol",
			Brand: "Volkswagen",
			Year:  2010,
		},
	}

	fmt.Println(myCar)

	// Using interface methods
	fmt.Println(StartVehicle(myCar))
	fmt.Println(RunVehicle(myCar))
}