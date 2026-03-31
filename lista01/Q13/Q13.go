package main

import "fmt"

var nota float64

var tipo string

func main() {

	for {

		fmt.Println("Digite a nota: ")

		fmt.Scan(&nota)

		if nota >= 9 && nota >= 10 {

			tipo = "A"

		} else if nota >= 7.5 && nota < 9 {

			tipo = "B"

		} else if nota >= 6 && nota < 7.5 {

			tipo = "C"

		} else {

			tipo = "D"

		}

		if -1 < nota && nota < 11 {

			fmt.Printf("NOTA = %.1f, CONCEITO = %v\n", nota, tipo)

			break

		} else {

			fmt.Println("INVALIDO")

			continue

		}

	}

}
