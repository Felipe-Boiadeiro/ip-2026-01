package main
import "fmt"

func main() {
	var n int
	fmt.Println("Digite um número: ")
	fmt.Scan(&n)

	if n < 0 {
		fmt.Println("O número é negativo")
	}else if n > 0 {
		fmt.Println("O número é positivo")
	} else {
		fmt.Println("O número é nulo")
	}
}
