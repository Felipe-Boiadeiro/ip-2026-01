package main

import "fmt"

func main() {
	n, soma := 0, 0

	fmt.Scan(&n)

	for i := 0; i <= n; i++ {

		soma += i

	}

	fmt.Println(soma)

}
