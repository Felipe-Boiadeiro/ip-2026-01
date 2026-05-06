package main

import "fmt"

func potencia(x, y int) int {
	valor := 1
	for i := 1; i <= y; i++ {
		valor *= x

	}
	return valor

}

func main() {
	x, n := 0, 0
	fmt.Scan(&x)
	fmt.Scan(&n)
	fmt.Println(potencia(x, n))

}
