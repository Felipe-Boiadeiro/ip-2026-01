package main

import "fmt"

func main() {
	n1, n2, n3 := 0, 0, 0
	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&n1)
	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&n2)
	fmt.Print("Digite o terceiro número: ")
	fmt.Scan(&n3)

	m := maior(n1, n2, n3)
	fmt.Printf("O maior número é %d\n", m)

}
func maior(x, y, z int) int {
	if x >= y && y >= z {
		return x
	} else if y >= z {
		return y
	} else {
		return z
	}
}
