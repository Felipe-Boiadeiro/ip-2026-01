package main

import "fmt"

func main() {

	var conta, tipo string
	var vol, valor float64
	fmt.Print("Informe o tipo de conta \033[90mR = Residencial, C = comercial, I= industrial\033[0m\n: ")
	fmt.Scan(&conta)

	fmt.Print("Informe a quantidade de água em metros cúbicos: ")
	fmt.Scan(&vol)

	switch conta {
	case "R", "r":
		valor = 5 + 0.05*vol
		tipo = "residencial"

	case "C", "c":
		tipo = "comercial"
		if vol <= 80 {
			valor = 500
		} else {
			valor = 500 + 0.25*(vol-80)
		}

	case "I", "i":
		tipo = "industrial"
		if vol <= 100 {
			valor = 800
		} else {
			valor = 800 + 0.04*(vol-100)
		}
	}
	fmt.Printf("Sua conta é %v e o valor a ser pago é R$ %.2f\n", tipo, valor)
}

