package main

import (
	"fmt"
	"time"
)

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// using var
	var name string = "Muttakin Islam Hasib"
	var age int = time.Now().Year() - 1999
	var isCool bool = true

	// shorthand
	email := "hi@muttakin.me"

	fmt.Println(name, age, email)
	// get the type name
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)
}
