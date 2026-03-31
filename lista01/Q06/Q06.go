package main

import "fmt"

var n int

var celsius float64

func main() {

	fmt.Println("Quantas operações de conversão de temperatura você deseja realizar: ")

	fmt.Scan(&n)

	tempf := make([]float64, n)

	if n == 1 {

		fmt.Println("Informe a temperatura a ser convertida: ")

	} else {

		fmt.Println("Informe as tempeaturas a serem convertidas: ")

	}

	for i := 0; i < n; i++ {

		fmt.Scan(&tempf[i])

	}

	for i := 0; i < n; i++ {

		celsius = 5.0 * (tempf[i] - 32) / 9.0

		fmt.Printf("%v FAHRENHEIT EQUIVALE A %.2f CELSIUS\n", tempf[i], celsius)

	}

}
