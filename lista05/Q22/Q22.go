package main

import "fmt"

func main() {
	conta := map[int]float64{}
	operacao := 0

	for i := 1; i <= 10; i++ {
		codigo := 0
		saldo := 0.0

		fmt.Printf("Informe o codigo da %dª conta: ", i)
		fmt.Scan(&codigo)

		fmt.Printf("Informe o saldo da %dª conta: ", i)
		fmt.Scan(&saldo)

		conta[codigo] = saldo
	}

	for {
		fmt.Println("\n===== MENU =====")
		fmt.Println("1 - Efetuar depósito")
		fmt.Println("2 - Efetuar saque")
		fmt.Println("3 - Consultar o ativo bancário")
		fmt.Println("4 - Finalizar o programa")
		fmt.Println("================")
		fmt.Print("Opção: ")
		fmt.Scan(&operacao)

		switch operacao {

		case 1:
			codigo := 0
			fmt.Print("Informe o código da conta: ")
			fmt.Scan(&codigo)

			saldoAtual, ok := conta[codigo]

			if ok {
				deposito := 0.0
				fmt.Print("Informe o valor do depósito: ")
				fmt.Scan(&deposito)

				conta[codigo] = saldoAtual + deposito
				fmt.Println("Novo saldo:", conta[codigo])
			} else {
				fmt.Println("Conta não encontrada")
			}

		case 2:
			codigo := 0
			fmt.Print("Informe o código da conta: ")
			fmt.Scan(&codigo)

			saldoAtual, ok := conta[codigo]

			if ok {
				saque := 0.0
				fmt.Print("Informe o valor do saque: ")
				fmt.Scan(&saque)

				if saque > saldoAtual {
					fmt.Println("Saldo insuficiente")
				} else {
					conta[codigo] = saldoAtual - saque
					fmt.Println("Novo saldo:", conta[codigo])
				}
			} else {
				fmt.Println("Conta não encontrada")
			}

		case 3:
			total := 0.0
			for _, saldo := range conta {
				total += saldo
			}
			fmt.Println("Ativo bancário:", total)

		case 4:
			fmt.Println("Encerrando...")
			return

		default:
			fmt.Println("Opção inválida")
		}
	}
}
