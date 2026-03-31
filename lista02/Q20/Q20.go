package main

import (
	"fmt"
)

func main() {
	var tipo int
	var val, valFinal float64
	fmt.Print("Informe a condição de pagamento: ")
	fmt.Scan(&tipo)

	switch {
	case tipo == 1:
		valFinal = val * 0.9

	case tipo == 2:
		valFinal = val * 0.95

	case tipo == 3:
		valFinal = val

	case tipo == 4:
		valFinal = val * 1.1
	default:
		fmt.Println("Condição inválida")
		return
	}
	fmt.Printf("O valor a ser pago é R$ %.2f", valFinal)
}
