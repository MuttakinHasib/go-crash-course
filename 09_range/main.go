package main

import "fmt"

func main() {
	ids := []int{1, 2, 3}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Println("ID: ", id)
	}

	emails := map[string]string{"hasib": "hi@muttakin.me", "sabit": "hi@sabit.me"}

	for k, v := range emails {
		fmt.Printf("%s:%s\n", k, v)
	}
}
