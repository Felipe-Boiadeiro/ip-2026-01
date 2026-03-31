package main

import "fmt"

func main() {
	var idade int

	fmt.Print("Informe a idade: ")
	fmt.Scan(&idade)

	switch {
	case idade >= 0 && idade <= 2:
		fmt.Println("Recém-nascido")

	case idade >= 3 && idade <= 11:
		fmt.Println("Criança")

	case idade >= 12 && idade <= 19:
		fmt.Println("Adolescente")

	case idade >= 20 && idade <= 55:
		fmt.Println("Adulto")

	case idade < 0:
		fmt.Println("Inválido")

	default:
		fmt.Println("Idoso")
	}

}
