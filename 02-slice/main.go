package main

import "fmt"

func main() {
	// Slices are references to arrays

	// Since slices are passed by reference,
	// they point to an memory address where the array is stored
	// If you modify the slice, you are modifying the array
	// this goes for copies of the slice and functions as well

	// If an array surpasses the capacity of a slice,
	// a new array is created and the values are copied over
	// therefore the slice will point to a new memory address
	// which means that if you modify the slice.
	// you are now modifying a new array

	// Initialize a slice with 5 int elements
	slice := make([]int, 5)
	
	// Assign values to the slice
	slice[0] = 1

	fmt.Println()
	fmt.Println("Slice1   |",slice)
	fmt.Println("-------- | --------")
	fmt.Println("Size     |",len(slice))
	fmt.Println("Capacity |",cap(slice))
	fmt.Println()

	// Initialize a slice with 5 int elements and capacity of 10
	slice2 := make([]int, 2, 4)

	fmt.Println("Slice2   |",slice2)
	fmt.Println("-------- | --------")
	fmt.Println("Size     |",len(slice2))
	fmt.Println("Capacity |",cap(slice2))
	fmt.Println()

	// When the capacity of a slice is reached, a new array is created and the values are copied over
	
	// Append values to slice2
	slice2 = append(slice2, 1, 2, 3)

	fmt.Println("After appending 3 values to slice2:")
	fmt.Println()
	fmt.Println("Slice2   |",slice2)
	fmt.Println("-------- | --------")
	fmt.Println("Size     |",len(slice2))
	fmt.Println("Capacity |",cap(slice2),"<--")

	// Showin that a slice always doubles the capacity when the capacity is reached
	fmt.Println()
	fmt.Println("Slice2 length and capacity appending values: ")
	fmt.Println()
	fmt.Println("Size | Capacity")
	fmt.Println("---- | --------")

	oldCap := cap(slice2)
	fmt.Println(len(slice2), "   |", cap(slice2),"<--")
	for i := 0; i < 20; i++ {
		slice2 = append(slice2, i)
		if newcap := cap(slice2); newcap != oldCap{
			oldCap = newcap
			if len(slice2) < 10 {
				fmt.Println(len(slice2), "   |", cap(slice2),"<--")
			} else {
				fmt.Println(len(slice2), "  |", cap(slice2),"<--")
			}
		}
	}
}