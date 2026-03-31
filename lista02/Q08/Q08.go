package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite um número: ")
	fmt.Scan(&n)

	if n > 20 && n < 90 {
		fmt.Println("O número está compreendido entre 20 e 90")
	} else {
		fmt.Println("O número não está compreendido entre 20 e 90")
	}

}
