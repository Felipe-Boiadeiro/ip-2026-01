package main

import "fmt"

func main() {

	n1, n2, multiplicacao := 0, 0, 0
	fmt.Scan(&n1)
	fmt.Scan(&n2)

	for i := 1; i <= n2; i++ {

		multiplicacao += n1

	}
	fmt.Println(multiplicacao)
}
