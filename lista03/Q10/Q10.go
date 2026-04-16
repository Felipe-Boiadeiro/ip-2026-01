package main

import "fmt"

func main() {

	n := 0
	fibonacii := []int{0, 1, 1}
	fmt.Scan(&n)

	for i := 2; i <= n; i++ {
		if n < 3 {
			break
		}

		fibonacii = append(fibonacii, fibonacii[i]+fibonacii[i-1])

	}
	for x := range fibonacii {
		fmt.Printf("%d ", fibonacii[x])
	}
}
