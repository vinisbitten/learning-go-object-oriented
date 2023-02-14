package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["a"] = 10
	m["b"] = 20
	m["c"] = 30
	fmt.Println(m)

	delete(m, "c")
	println(m["c"])

	_, exists := m["c"]
	println(exists)

	if value, exists := m["a"]; exists {
		println(value)
	}
}
