package main

import (
	"fmt"
)

func main() {
	var soma uint64
	var valor uint64 = 1

	for i := 0; i <= 64; i++ {
		soma += valor
		valor *= 2
	}
	fmt.Println(soma)

}
