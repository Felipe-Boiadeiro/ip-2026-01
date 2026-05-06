package main

import "fmt"

func main() {
	sequencia := []int{}
	sequencia2 := []int{}

	n := 0

	for i := 1; i <= 30; i++ {
		fmt.Scan(&n)
		sequencia = append(sequencia, n)
	}

	for i := range sequencia {
		if i%2 == 0 {
			sequencia2 = append(sequencia2, sequencia[i]*2)

		} else {
			sequencia2 = append(sequencia2, sequencia[i]*3)
		}
	}
	fmt.Println(sequencia2)
}
