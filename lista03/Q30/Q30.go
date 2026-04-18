package main

import (
	"fmt"
	"math"
)

func main() {

	r := .0

	fmt.Println("Informe o valora do raio em cm (deve estar entre 0 a 20cm, obedecendo um período de 0,5cm): ")
	fmt.Scan(&r)

	if r > 0 && r <= 20 {
		if r*2 == float64(int(r*2)) {
			volume := 4 * math.Pi * math.Pow(r, 3) / 3

			fmt.Printf("%.2f cm^3", volume/1000)

		}

	}

}
