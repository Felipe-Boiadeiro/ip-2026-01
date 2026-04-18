package main

import (
	"fmt"
	"math"
)

func main() {
	x, cosseno := .0, 1.0

	fmt.Scan(&x)

	termo := 1.0
	for i := 1; i <= 20; i++ {
		termo *= -(x * x) / float64((2*i-1)*(2*i))
		cosseno += termo
	}

	diferenca := cosseno - math.Cos(x)

	fmt.Println(cosseno)
	fmt.Println(diferenca)
}
