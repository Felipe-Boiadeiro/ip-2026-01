package main

import "fmt"

func main() {
	sequencia := []int{}
	frequencia := map[int]int{}
	n := 0
	for i := 1; i <= 50; i++ {
		fmt.Scan(&n)
		sequencia = append(sequencia, n)

	}
	for i := range sequencia {
		frequencia[sequencia[i]]++
	}
	moda := 0
	maiorFre := 0

	for chave, freq := range frequencia {
		if freq > maiorFre {
			maiorFre = freq
			moda = chave
		}
	}
	fmt.Println(moda)
}
