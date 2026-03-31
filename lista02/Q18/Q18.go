package main

import "fmt"

func main() {
	var dia int
	var valor1, valor2, preco float64
	var tipo string
	fmt.Print("Informe o dia da semana \033[90m1 = Domingo, 2 = Segunda, 3 = Terça, 4 = Quarta, 5 = Quinta, 6 = Sexta, 7 = Sábado\033[0m\n: ")
	fmt.Scan(&dia)
	fmt.Print("Informe o valor do DVD: ")
	fmt.Scan(&valor1)
	fmt.Print("Tipo de DVD \033[90mC = Comum, L = Lançamento\033[0m\n: ")
	fmt.Scan(&tipo)

	switch tipo {
	case "C", "c":
		valor2 = valor1

	case "L", "l":

		valor2 = valor1 * 1.15

	default:
		fmt.Println("Inválido")

	}
	if dia == 2 || dia == 3 || dia == 5 {
		preco = valor2 * 0.6
	} else if dia == 1 || dia == 4 || dia == 6 || dia == 7 {
		preco = valor2
	} else {
		fmt.Println("Inválido")
		return
	}
	fmt.Printf("O preço da locação será R$ %.2f", preco)

}
