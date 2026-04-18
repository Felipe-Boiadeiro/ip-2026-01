package main

import (
	"fmt"
	"math"
)

func main() {

	n := .0
	fmt.Scan(&n)

	sequencia := []int{}

	for i := 1; i <= n; i++ {

		x := i
		if n%math.Pow(10, float64(x)) != 0 {
			sequencia = append(sequencia, int(n%10))
		}
	}

}
