package main

import "fmt"

func main() {
	idades := []int{}
	alturas := []float64{}
	pesos := []float64{}

	idade, altura, peso := 0, .0, .0
	contador, loop, soma, media, magros := 0, 0, .0, 0, 0
	for {
		fmt.Print("Informe a idade: ")
		fmt.Scan(&idade)

		idades = append(idades, idade)
		fmt.Print("Informe a altura: ")
		fmt.Scan(&altura)

		alturas = append(alturas, altura)
		fmt.Print("Informe o peso: ")
		fmt.Scan(&peso)

		pesos = append(pesos, peso)

		fmt.Print("Caso queira adicionar mais pessoas digite 1: ")
		fmt.Scan(&loop)

		if loop != 1 {
			break
		}
	}

	for i := range idades {
		if idades[i] >= 50 {
			contador++
		} else if idades[i] >= 10 && idades[i] <= 20 {

			soma += alturas[i]
			media++

		}

	}
	for i := range pesos {
		if pesos[i] < 40 {
			magros++
		}
	}
	porcentagem := (float64(magros) / float64(len(pesos))) * 100

	fmt.Println("Quantidade de pessoas com 50 anos ou mais:", contador)
	fmt.Printf("Média das alturas de pessoas entre 10 e 20 anos: %.2f", soma/float64(media))
	fmt.Printf("Porcentagem de pessoas com peso inferior a 40 Kg: %.2f%%\n", porcentagem)
}
