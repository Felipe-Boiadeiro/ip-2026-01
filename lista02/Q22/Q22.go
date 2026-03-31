package main

import "fmt"

func main() {
	var salMin float64 = 788.00
	var vHora_extra float64 = 10.00
	var matricula int
	var hora_extra, inss, imposto, salLiq, salBruto float64
	fmt.Print("Informe a sua matrícula: ")
	fmt.Scan(&matricula)

	fmt.Print("Informe a quantidade de horas extras trabalhadas: ")
	fmt.Scan(&hora_extra)

	salBruto = 3*salMin + vHora_extra*hora_extra

	inss = 0
	imposto = 0

	if salBruto > 1500 {
		inss = 0.12 * salBruto
	}
	if salBruto > 2000 {
		imposto = 0.2 * salBruto
	}

	salLiq = salBruto - inss - imposto

	fmt.Printf("O salário líquido de %v é R$ %.2f\n", matricula, salLiq)

}
