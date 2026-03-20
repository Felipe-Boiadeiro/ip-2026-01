package main

import "fmt"

var sMinimo, qKw float64

var vlrKw, ccheio, cdesconto float64

func main() {

	fmt.Println("Informe o salário mínimo: ")

	fmt.Scan(&sMinimo)

	fmt.Println("Informe a quantidade de Kw gastos pela residência: ")

	fmt.Scan(&qKw)

	//calculo

	vlrKw = sMinimo * 0.007

	ccheio = vlrKw * qKw

	cdesconto = ccheio * 0.9

	fmt.Printf("Custo por kW: R$ %.2f\n", vlrKw)

	fmt.Printf("Custo do consumo: R$ %.2f\n", ccheio)

	fmt.Printf("Custo com desconto: R$ %.2f\n", cdesconto)

}
