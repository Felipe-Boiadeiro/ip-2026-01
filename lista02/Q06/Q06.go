package main
import "fmt"

func main() {
	var a, b int
	fmt.Print("Digite um número A: ")
	fmt.Scan(&a)

	fmt.Print("Digite um número B: ")
	fmt.Scan(&b)

	if a % b == 0 {
		fmt.Println("B é um divisor de A")
	} else {
		fmt.Println("B não é um divisor de A")
	}
}
