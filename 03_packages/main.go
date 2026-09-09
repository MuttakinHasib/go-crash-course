package main

import (
	"fmt"
	"math"

	"github.com/MuttakinHasib/go-crash-course/03_packages/str_util"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(2.7))
	fmt.Println(str_util.ReverseString("Hello"))
}
