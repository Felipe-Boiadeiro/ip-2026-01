package main

import "fmt"

var conta int

var qagua, custo float64

var tipo string

func main() {

	fmt.Println("Informe sua conta, consumo de água em metros cúbicos e o tipo de consumidor: ")

	fmt.Scan(&conta, &qagua, &tipo)

	switch tipo {

	case "r", "R":

		custo = 5 + qagua*0.05

		fmt.Println(custo)

	case "i", "I":

		if qagua <= 100 {

			custo = 800

		} else {

			custo = 800 + (qagua-100)*0.04

		}

	case "c", "C":

		if qagua <= 80 {

			custo = 500

		} else {

			custo = 500 + (qagua-80)*0.25

		}

	}

	fmt.Println("CONTA = ", conta)

	fmt.Printf("VALOR DA CONTA = %.2f", custo)

}
