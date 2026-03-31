package main

import (
	"fmt"
)

var a1, r, n, soma int

func main() {

	fmt.Println("Digite o primeiro, a razão e a quantidade de elementos da progressão: ")

	fmt.Scan(&a1, &r, &n)

	soma = (a1 + (a1 + r*(n-1))) * n / 2

	fmt.Println(soma)

}
