package main

import "fmt"

func main() {
	// Initialize an array with 10 int elements
	var x [10]int

	fmt.Println()
	fmt.Println("Assigning values for x")
	fmt.Println()
	fmt.Println("Array x:",x)
	
	// Assign values to the array
	x[0] = 5
	x[1] = 7
	x[2] = 2

	fmt.Println("Array x after -->",x)
	fmt.Println()

	// Trying to assign a value to an index that is out of range will cause a runtime error
	// x[10] = 100
	
	// Initialize an array with 5 int elements and assign values
	y := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Initializing y with values")
	fmt.Println()
	fmt.Println("Array y:",y)
	fmt.Println()

	z := [4]string{"Hello", "World", "Vinícius", "Bittencourt"}
	// Getting elements from an array
	fmt.Println("Getting elements from z")
	fmt.Println()
	fmt.Println("Array z: ",z)
	fmt.Println()
	fmt.Println("z[0]  :",z[0]) // Hello
	fmt.Println("z[:2] :",z[:2]) // Hello World
	fmt.Println("z[2:] :",z[2:]) // Vinícius Bittencourt
}