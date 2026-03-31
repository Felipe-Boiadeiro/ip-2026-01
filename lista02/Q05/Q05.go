package main
import "fmt"

func main() {
	var n int
	fmt.Println("Digite um número: ")
	fmt.Scan(&n)

	if n % 5 == 0 {
		fmt.Println("O número é divisível por 5")
	} else {
		fmt.Println("O número não é divisível por 5")
	}
}
