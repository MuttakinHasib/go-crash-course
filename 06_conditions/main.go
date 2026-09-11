package main

import "fmt"

func main() {
	x := 5
	color := "red"

	if x >= 2 {
		fmt.Printf("%d is grater then 2", x)
	}

	switch color {
	case "red":
		fmt.Println("Color is Red")
	default:
		fmt.Println("Color is ", color)

	}

}
