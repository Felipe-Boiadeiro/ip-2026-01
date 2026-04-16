package main

import "fmt"

func fatorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * fatorial(x-1)
}

func main() {
	y := .0
	fmt.Scan(&y)

	soma := y

	for i := 1; i < 20; i++ {
		termo := y / float64(fatorial(i))

		if i%2 == 1 {
			soma -= termo
		} else {
			soma += termo
		}
	}

	fmt.Println(soma)
}
