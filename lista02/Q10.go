package main

import "fmt"

func main() {
	var destino, tipo int
	fmt.Print("Informe a região de destino: ")
	fmt.Scan(&destino)

	switch destino {
	case (1):
		fmt.Print("A viagem inclui retorno: ")
		fmt.Scan(&tipo)

		switch tipo {
		case (1):
			fmt.Println("O valor é R$ 900,00")
		case (2):
			fmt.Println("O valor é R$ 500,00")
		default:
			fmt.Println("Valor inválido")
		}

	case (2):
		fmt.Print("A viagem inclui retorno: ")
		fmt.Scan(&tipo)

		switch tipo {
		case (1):
			fmt.Println("O valor é R$ 650,00")
		case (2):
			fmt.Println("O valor é R$ 350,00")
		default:
			fmt.Println("Valor inválido")
		}
	case (3):
		fmt.Print("A viagem inclui retorno: ")
		fmt.Scan(&tipo)

		switch tipo {
		case (1):
			fmt.Println("O valor é R$ 600,00")
		case (2):
			fmt.Println("O valor é R$ 350,00")
		default:
			fmt.Println("Valor inválido")
		}
	case (4):
		fmt.Print("A viagem inclui retorno: ")
		fmt.Scan(&tipo)

		switch tipo {
		case (1):
			fmt.Println("O valor é R$ 550,00")
		case (2):
			fmt.Println("O valor é R$ 300,00")
		default:
			fmt.Println("Valor inválido")

		}
	default:
		fmt.Println("Valor inválido")

	}
}
