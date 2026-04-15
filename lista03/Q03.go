package main

import "fmt"

func main() {

	carlos := .0
	joao := .0
	meses := 0

	fmt.Print("Informe o salário de Carlos: ")
	fmt.Scan(&carlos)

	joao = carlos / 3

	for {
		joao = joao * 1.05
		carlos = carlos * 1.02
		meses++

		if joao > carlos {
			fmt.Printf("Vai demorar %d meses para João ter mais que carlos", meses)
			break

		}

	}

}

