package main
import "fmt"

func main() {
	var n int64
	fmt.Println("Informe o número para verificar se ele é par ou ímpar: ")
	fmt.Scan(&n)

	if n % 2 == 0 {
		fmt.Println("O número é par")

	} else {
		fmt.Println("O número é ímpar")
	}
}
