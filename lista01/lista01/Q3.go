package main

import (
	"fmt"

	"os"
)

func main() {

	fmt.Println("Digite a centena: ")

	var cent int

	fmt.Scan(&cent)

	if cent <= 0 || cent > 9 {

		fmt.Println("Digito inválido")

		os.Exit(1)

	}

	fmt.Println("Digite a dezena: ")

	var dez int

	fmt.Scan(&dez)

	if dez < 0 || dez > 9 {

		fmt.Println("Digito inválido")

		os.Exit(1)

	}

	fmt.Println("Digite a unidade: ")

	var uni int

	fmt.Scan(&uni)

	if uni < 0 || uni > 9 {

		fmt.Println("Digito inválido")

		os.Exit(1)

	}

	concat := cent*100 + dez*10 + uni

	sqr := concat * concat

	fmt.Println(concat, "-", sqr)

}
