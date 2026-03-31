package main

import "fmt"

var r, h, at, custo float64

func main() {

	pi := 3.14159

	fmt.Println("Informe o raio da lata: ")

	fmt.Scan(&r)

	fmt.Println("Informe a altura da lata")

	fmt.Scan(&h)

	at = (2 * pi * r * r) + (2 * pi * r * h)

	custo = at * 100

	fmt.Printf("O VALOR DO CUSTO E = %.2f", custo)

}
