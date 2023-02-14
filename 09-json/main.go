package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	// binding with json tags
	Name string `json:"name"`
	Age  int `json:"age"`
}

func main() {
	// Creating a "json" string and unmarshalling it into a struct
	jsonString := `{"name":"John", "age": 30}`

	var john Person

	json.Unmarshal([]byte(jsonString), &john)

	fmt.Println(john)

	// Creating a struct and marshalling it into a "json" string
	jane := Person{
		Name: "Jane",
		Age:  25,
	}

	jsonJane, _ := json.Marshal(jane)

	fmt.Println(string(jsonJane))
}