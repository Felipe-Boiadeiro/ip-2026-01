package main

import (
	"fmt"
)

var x, y int

func main() {

	fmt.Println("Digite o número e a quantidade de sequência de pares desejados a ver: ")

	fmt.Scan(&x, &y)

	if x%2 != 0 {

		fmt.Println("O PRIMEIRO NUMERO NAO E PAR")

		return

	}

	for i := 0; i < y; i++ {

		if i > 0 {

			fmt.Print(" ")

		}

		fmt.Print(x + 2*i)

	}

}
