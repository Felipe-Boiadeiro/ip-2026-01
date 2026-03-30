package main

import "fmt"

func main() {
	var preco_inicial, valor_ar, valor_pintura, valor_vidro, valor_direcao, valor_final float64
	var ar, pint, vidro, direc int

	fmt.Print("Informe o valor inicial do carro: ")
	fmt.Scan(&preco_inicial)

	fmt.Print("Você deseja adicionar Ar condicionado (1 para sim, 0 para não): ")
	fmt.Scan(&ar)

	if ar == 1 {
		valor_ar = 1750
	} else if ar == 0 {
		valor_ar = 0
	} else {
		fmt.Println("Reposta inválida")
		return
	}

	fmt.Print("Você deseja adicionar Pintura metálica (1 para sim, 0 para não): ")
	fmt.Scan(&pint)

	if pint == 1 {
		valor_pintura = 800
	} else if pint == 0 {
		valor_pintura = 0
	} else {
		fmt.Println("Reposta inválida")
		return
	}

	fmt.Print("Você deseja adicionar Vidro elétrico (1 para sim, 0 para não): ")
	fmt.Scan(&vidro)

	if vidro == 1 {
		valor_vidro = 1200
	} else if vidro == 0 {
		valor_vidro = 0
	} else {
		fmt.Println("Reposta inválida")
		return
	}

	fmt.Print("Você deseja adicionar Direção hidráulica (1 para sim, 0 para não): ")
	fmt.Scan(&direc)

	if direc == 1 {
		valor_direcao = 2000
	} else if direc == 0 {
		valor_direcao = 0
	} else {
		fmt.Println("Reposta inválida")
		return
	}

	valor_final = preco_inicial + valor_ar + valor_pintura + valor_vidro + valor_direcao

	fmt.Printf("O valor do carro será %.2f", valor_final)
}
