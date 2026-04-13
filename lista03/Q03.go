package main

import "fmt"

func main() {
	var sal_carlos, sal_joao float64
	var meses int
	fmt.Print("Informe o salário de Carlos: ")
	fmt.Scan(&sal_carlos)

	sal_joao = sal_carlos / 3

	for {
		sal_joao = sal_joao * 1.05
		sal_carlos = sal_carlos * 1.02
		meses++
		if sal_joao >= sal_carlos {
			fmt.Printf("Para que o valor pertencente a João iguale ou ultrapasse o valor pertencente a Carlos, passarão %v meses", meses)
			break
		}

	}
}
