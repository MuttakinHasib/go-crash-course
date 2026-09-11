package main

import "fmt"

func greeting(name string) string {
	return "Assalamualaikum, " + name
}

func getSum(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(greeting("Hasib"))
	a, b := 1, 2
	fmt.Printf("Summation if %d + %d = %d\n", a, b, getSum(a, b))
}
