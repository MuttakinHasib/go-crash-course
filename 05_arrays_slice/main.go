package main

import "fmt"

func main() {

	// Arrays
	var fruits [2]string

	// Assign values
	fruits[0] = "Apple"
	fruits[1] = "Orange"

	// Declare and assign
	foods := [2]string{"Rice", "Vegetable"}

	fmt.Println(fruits)
	fmt.Println(fruits[1])
	fmt.Println(foods)

	// Slice
	flowers := []string{"Rose", "Sun Flower"}

	flowersCount := len(flowers)

	fmt.Println(flowers)
	fmt.Println(flowersCount)

}
