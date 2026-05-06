package main

import "fmt"

func main() {
	sequencia := []int{}

	for i := 100; i >= 1; i-- {
		sequencia = append(sequencia, i)
	}
	fmt.Println(sequencia)
}
