package main

import "fmt"

func main() {

	var sal, nsal float64

	fmt.Println("Informe o salário: ")

	fmt.Scan(&sal)

	if sal < 0 {

		fmt.Println("Valor inválido")

		return

	}

	if sal <= 300 {

		nsal = sal * 1.5

	} else {

		nsal = sal * 1.3

	}

	fmt.Printf("SALARIO COM REAJUSTE = %.2f\n", nsal)

}
