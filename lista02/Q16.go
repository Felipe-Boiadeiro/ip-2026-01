package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, delta, x1, x2 float64
	fmt.Print("Digite os coeficientes da equação de segundo grau: ")
	fmt.Scan(&a, &b, &c)

	//Delta
	delta = b*b - 4*a*c
	x1 = (-b + math.Pow(delta, 0.5)) / 2 * a
	x1 = (-b - math.Pow(delta, 0.5)) / 2 * a

	if delta > 0 {
		fmt.Println("A função possui duas raízes distintas")
		fmt.Printf("x'= %.2f\n x''= %.2f\n", x1, x2)

	} else if delta == 0 {
		fmt.Println("A função possui duas raízes iguais")
		fmt.Printf("x' e x'' = %.2f", x1)
	} else {
		fmt.Println("A função possui raízes imaginárias")
		fmt.Printf("x'= %.2f\n x''= %.2f\n", x1, x2)
	}

}
