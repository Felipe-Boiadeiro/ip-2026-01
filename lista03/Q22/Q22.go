package main

import "fmt"

func main() {
	soma := .0

	for i := 1; i <= 37; i++ {
		soma += float64((38-i)*(39-i)) / float64(i)
	}

	fmt.Println(soma)
}
