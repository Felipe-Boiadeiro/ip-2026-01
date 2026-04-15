package main

import "fmt"

func main() {
	var n int

	for {
		fmt.Print("Digite um número: ")
		fmt.Scan(&n)

		if n > 0 {
			i := 1

			for ; i*i <= n; i++ {
				if i*i == n {
					fmt.Printf("%v é um quadrado perfeito\n", n)
					break
				}
			}

			// Se saiu do loop porque i*i passou de n, então não encontrou
			if i*i > n {
				fmt.Printf("%v não é um quadrado perfeito\n", n)
			}

		} else {
			fmt.Println("Parando")
			break
		}
	}
}
