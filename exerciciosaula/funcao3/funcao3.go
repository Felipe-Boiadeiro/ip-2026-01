package main

import "fmt"

func main() {
	n1, n2, n3 := 0.0, 0.0, 0.0
	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&n1)
	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&n2)
	fmt.Print("Digite o terceiro número: ")
	fmt.Scan(&n3)
	m := media(n1, n2, n3)

	fmt.Printf("A média é %.2f\n", m)

}
func media(x, y, z float64) float64 {
	return (x + y + z) / 3
}
