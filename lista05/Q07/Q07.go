package main

import "fmt"

func main() {
	sequencia := []int{}
	i := 1
	for {
		if i%2 != 0 {
			sequencia = append(sequencia, i)
		}
		i++

		if len(sequencia) == 100 {
			break
		}

	}
	fmt.Println(sequencia)
}
