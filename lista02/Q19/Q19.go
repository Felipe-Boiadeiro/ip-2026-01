package main

import (
	"fmt"
	"math"
)

func main() {
	var tipo int
	var r, h, volume, area float64
	fmt.Print("Informe a figura desejada \033[90m(1-cone / 2-cilindro /3-esfera)\033[0m\n:")
	fmt.Scan(&tipo)

	if tipo >= 1 && tipo <= 3 {
		fmt.Print("Informe o raio: ")
		fmt.Scan(&r)
		fmt.Print("Informe a altura: ")
		fmt.Scan(&h)
	} else {
		return
	}

	switch {
	case tipo == 1:
		volume = 3.14 * r * r * h / 3
		area = 3.14 * r * math.Pow(r*r+h*h, 0.5)

	case tipo == 2:
		volume = 3.14 * r * r * h
		area = 2 * 3.14 * r * h

	case tipo == 3:
		volume = 4 * 3.14 * r * r * r / 3
		area = 4 * 3.14 * r * r
	}

	fmt.Printf("O volume é %f m³ e a área é %f m²", volume, area)
}
