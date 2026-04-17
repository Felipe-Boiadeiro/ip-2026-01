package main

import "fmt"

func main() {

	b, n, valor := 0, 0, 1
	fmt.Scan(&b)
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		valor *= b
	}

	fmt.Println(valor)
}
