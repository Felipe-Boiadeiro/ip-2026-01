package main

import (
	"fmt"
	"sort"
)

func main() {
	sequencia := []int{}
	n := 0
	for i := 1; i <= 10; i++ {
		fmt.Scan(&n)
		sequencia = append(sequencia, n)
	}
	sort.Ints(sequencia)
	fmt.Println(sequencia)
}
