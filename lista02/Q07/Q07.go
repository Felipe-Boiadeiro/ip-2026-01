package main

import "fmt"

func main() {

	var a, b, c, MENOR, INTER, MAIOR int
	fmt.Print("Informe 3 números distintos: ")
	fmt.Scan(&a, &b, &c)

	if a < b && a < c {
		MENOR = a
		if b < c {
			INTER = b
			MAIOR = c
		} else {
			INTER = c
			MAIOR = b
		}

	} else if b < a && b < c {
		MENOR = b
		if a < c {
			INTER = a
			MAIOR = c
		} else {
			INTER = c
			MAIOR = a
		}
	} else {
		MENOR = c
		if a < b {
			INTER = a
			MAIOR = b
		} else {
			INTER = b
			MAIOR = a
		}
	}

	fmt.Printf("%v, %v, %v\n", MENOR, INTER, MAIOR)

}
