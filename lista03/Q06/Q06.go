package main

import "fmt"

func main() {
	n, triangular := 0, false
	fmt.Scan(&n)
	msg := fmt.Sprintln(n, "não é um número triangular")

	for i := 1; i <= n; i++ {
		if i*(i+1)*(i+2) == n {
			fmt.Println(n, "é um número triangular")
			triangular = true
			break
		}

	}
	if triangular == false {
		fmt.Print(msg)
	}
}
