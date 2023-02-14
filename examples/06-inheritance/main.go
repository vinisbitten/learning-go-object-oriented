package main

import "fmt"

// Struct
type Car struct {
	Name string
	Brand string
	Year  int
}

// Methods
func (c Car) Run() string {
	return "Running..."
}

// Struct inheritance
type SuperCar struct {
	Car
	TopSpeed int
}

func main(){
	mySuperCar := SuperCar{
		Car: Car{
			Name: "Bugatti Veyron",
			Brand: "Bugatti",
			Year: 2010,
		},
		TopSpeed: 300,
	}

	fmt.Println(mySuperCar)
}