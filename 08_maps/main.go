package main

import "fmt"

func main() {
	// Define map

	emails := make(map[string]string)

	// Assign key value pair
	emails["Muttakin"] = "hi@muttakin.me"
	emails["Sabit"] = "hi@sabit.me"
	emails["Asraful"] = "hi@asraful.me"
	fmt.Println(emails)

	delete(emails, "Asraful")

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Muttakin"])
}
