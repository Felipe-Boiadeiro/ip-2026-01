package main

import "fmt"

func main() {
	var n int
	fmt.Print("Informe um número com 3 algarismos: ")
	fmt.Scan(&n)

	if n >= 100 && n <= 999 {

		dez := (n / 10) % 10
		fmt.Printf("O algarismo da casa das dezenas é %v\n", dez)
	} else {
		fmt.Println("O número deve ter 3 algarismos")
	}

}
