package main

import "fmt"

func main() {

	var n1, n2, n3, media float64

	fmt.Println("Digite as três notas separadas por espaço:")

	fmt.Scan(&n1, &n2, &n3)

	media = (n1 + n2 + n3) / 3

	fmt.Println("A média é:", media)

	if media >= 6 {

		fmt.Println("APROVADO")

	} else {

		fmt.Println("REPROVADO")

	}

}
