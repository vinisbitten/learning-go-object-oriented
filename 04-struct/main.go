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

func main(){
	myCar := Car{
		Name: "Gol",
		Brand: "Volkswagen",
		Year: 2010,
	}

	fmt.Println(myCar)
	fmt.Println(myCar.Run())
}