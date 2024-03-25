package main

import "fmt"

func main() {

	var (
		age    int
		weight float32
		name   string
	)

	age = 42
	weight = 90.4
	name = "Yakunin Vasily"

	fmt.Printf("Name: %s\nAge: %d\nWeight: %f", name, age, weight)

}
