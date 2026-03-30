package main

import "fmt"

func main() {
	var dia, mes, ano int
	var ext string

	fmt.Print("Informe a data desejada: ")
	fmt.Scan(&dia, &mes, &ano)

	if !(dia >= 1 && dia <= 31 && ano > 0 && mes >= 1 && mes <= 12) {
		fmt.Println("Inválido")
		return
	} else {
		switch {
		case mes == 1:
			ext = "Janeiro"

		case mes == 2:
			ext = "Fevereiro"

		case mes == 3:
			ext = "Março"

		case mes == 4:
			ext = "Abril"

		case mes == 5:
			ext = "Maio"

		case mes == 6:
			ext = "Junho"

		case mes == 7:
			ext = "Julho"

		case mes == 8:
			ext = "Agosto"

		case mes == 9:
			ext = "Setembro"

		case mes == 10:
			ext = "Outubro"

		case mes == 11:
			ext = "Novembro"

		case mes == 12:
			ext = "Dezembro"

		}

	}

	fmt.Printf("%v de %v de %v\n", dia, ext, ano)

}
