package main

import "fmt"

func main() {
	var idade int
	fmt.Print("Informe a sua idade: ")
	fmt.Scan(&idade)

	if idade < 16 && idade > -1 {
		fmt.Println("Não-eleitor")
	} else if idade >= 18 && idade <= 65 {
		fmt.Println("Eleitor obrigatório")
	} else if idade >= 16 && idade < 18 || idade > 65 {
		fmt.Println("Eleitor facultativo")
	} else {
		fmt.Println("Inválido")
	}
}
