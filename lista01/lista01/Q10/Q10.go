package main

import "fmt"

var a, b, c, d, det float64

func main() {

	fmt.Println("Informe o a11")

	fmt.Scan(&a)

	fmt.Println("Informe o a12")

	fmt.Scan(&b)

	fmt.Println("Informe o a21")

	fmt.Scan(&c)

	fmt.Println("Informe o a22")

	fmt.Scan(&d)

	det = a*d - b*c

	fmt.Printf("O VALOR DO DETERMINANTE E = %.2f", det)

}
