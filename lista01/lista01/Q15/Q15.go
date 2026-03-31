package main

import "fmt"

var n int

func main() {

	fmt.Println("Digite um número entre 5 e 2000: ")

	fmt.Scan(&n)

	for i := 1; i <= n; i++ {

		if i%2 == 0 {

			fmt.Printf("%d^2 = %d\n", i, i*i)

		}

	}

}
