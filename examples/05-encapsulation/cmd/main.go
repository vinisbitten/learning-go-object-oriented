package main

// Go provides encapsulation in package level
// So the lower case fields are not visible outside the package
// We will use modules to make them visible

import (
	Vehicle "encapsulation/pkg"
	"fmt"
)

func main() {
	// Create a new instance of the struct
	thisCar := Vehicle.Car{
		Name:  "Focus",
		Brand: "Ford",
		Year:  2018,
		// topSpeed: 200, --> unknown field
	}

	// Print the struct (will print topSpeed as 0)
	fmt.Println(thisCar)

	otherCar := Vehicle.ImportedCar()

	// Since otherCar is declared inside the other package 
	// Printing will print the topSpeed we set inside the package
	fmt.Println(otherCar)
}
