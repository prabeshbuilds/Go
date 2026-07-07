package main

import "fmt"

func main() {
	type Person struct {
		Name string
		Age  int
	}

	person := Person{
		Name: "Alice",
		Age:  25,
	}

	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
}