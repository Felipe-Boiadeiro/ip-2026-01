package main
import "fmt"

func main() {
	var n1, n2 int
	fmt.Println("Digite dois números: ")
	fmt.Scan(&n1, &n2)

	soma := n1 + n2

	if soma > 20 {
		fmt.Printf("A soma dos números somado com 8 é %v\n", soma + 8)
	} else {
		fmt.Printf("A soma dos números subtraido 5 é %v\n", soma - 5)
	}
}
