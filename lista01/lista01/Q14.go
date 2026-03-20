package main

import (
	"fmt"

	"math"
)

var h, a, v float64

func main() {

	fmt.Println("Informe a medida da altura e de uma aresta respectivamente: ")

	fmt.Scan(&h, &a)

	v = h * a * a * math.Sqrt(3) / 2

	fmt.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS", v)

}
