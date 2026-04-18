package main

import "fmt"

func main() {
	contador, n1, n2 := 0, 0, 0

	fmt.Scan(&n1)
	fmt.Scan(&n2)

	for n1 >= n2 {
		n1 -= n2
		contador++
	}

	fmt.Println("Quociente =", contador)
	fmt.Println("Resto =", n1)
}
