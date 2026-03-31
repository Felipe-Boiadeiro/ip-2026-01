package main

import "fmt"

func main() {
	var compra, venda float64
	fmt.Print("Informe o valor da compra: ")
	fmt.Scan(&compra)

	if compra < 10 && compra > 0 {
		venda = compra * 1.7
	} else if compra >= 10 && compra < 30 {
		venda = compra * 1.5
	} else if compra >= 30 && compra < 50 {
		venda = compra * 1.4
	} else if compra >= 50 {
		venda = compra * 1.3
	} else {
		fmt.Println("Valor inválido")
		return

	}

	fmt.Printf("O valor de venda será %.2f", venda)

}
