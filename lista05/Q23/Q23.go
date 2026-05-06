package main

import "fmt"

func main() {
	janela := make([]int, 24)
	corredor := make([]int, 24)

	opcao := 0

	for {
		fmt.Println("\n===== ÔNIBUS =====")
		fmt.Println("1 - Janela")
		fmt.Println("2 - Corredor")
		fmt.Println("3 - Sair")
		fmt.Print("Escolha: ")
		fmt.Scan(&opcao)

		if opcao == 3 {
			fmt.Println("Encerrando...")
			return
		}

		vetor := []int{}

		if opcao == 1 {
			vetor = janela
			fmt.Println("\nPoltronas disponíveis (Janela):")
		} else if opcao == 2 {
			vetor = corredor
			fmt.Println("\nPoltronas disponíveis (Corredor):")
		} else {
			fmt.Println("Opção inválida")
			continue
		}

		temLivre := false
		for i := 0; i < 24; i++ {
			if vetor[i] == 0 {
				fmt.Printf("%d ", i)
				temLivre = true
			}
		}

		if !temLivre {
			fmt.Println("\nNenhuma poltrona disponível nesse lado.")
			continue
		}

		num := 0
		fmt.Print("\nEscolha a poltrona: ")
		fmt.Scan(&num)

		if num < 0 || num >= 24 {
			fmt.Println("Poltrona inválida")
			continue
		}

		if vetor[num] == 1 {
			fmt.Println("Poltrona já ocupada")
		} else {
			vetor[num] = 1
			fmt.Println("Poltrona reservada com sucesso!")
		}

		if opcao == 1 {
			janela = vetor
		} else {
			corredor = vetor
		}

		total := 0
		for i := 0; i < 24; i++ {
			total += janela[i] + corredor[i]
		}

		if total == 48 {
			fmt.Println("\nÔnibus lotado!")
			return
		}
	}
}
