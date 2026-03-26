package main

import "fmt"

func main() {
	var x float64

	fmt.Print("Informe o valor de x: ")
	fmt.Scan(&x)

	if x == 2 {
		fmt.Println("f(x) = Indeterminado")
	} else {
		fmt.Printf("f(x) = %f", 8/(2-x))

	}

}
