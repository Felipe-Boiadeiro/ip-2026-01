package main

import (
	"fmt"
	"math"
)

func main() {

	a, sina := .0, .0
	fmt.Scan(&a)

	sina = a - (math.Pow(a, 3) / 6) + (math.Pow(a, 5) / 120) - (math.Pow(a, 7) / 5040)

	fmt.Println(sina)
}
