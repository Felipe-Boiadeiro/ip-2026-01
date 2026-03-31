package main

import "fmt"

var temp, pol, celsius, mm float64

func main() {

	fmt.Println("Informe a temperatura em fahrenheit a ser convertida: ")

	fmt.Scan(&temp)

	fmt.Println("Informe a quantidade de chuva em polegadas a ser convertida: ")

	fmt.Scan(&pol)

	celsius = (5*temp - 160) / 9

	mm = pol * 25.4

	fmt.Printf("O VALOR EM CELSIUS = %.2f\n", celsius)

	fmt.Printf("A QUANTIDADE DE CHUVA E = %.2f", mm)

}
