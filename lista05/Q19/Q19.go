package main

import "fmt"

func main() {
	vetor1 := []int{}
	vetor2 := []int{}

	num := 0
	divis := 0

	for i := 1; i <= 10; i++ {
		fmt.Scan(&num)
		vetor1 = append(vetor1, num)
	}

	for i := 1; i <= 5; i++ {
		fmt.Scan(&divis)
		vetor2 = append(vetor2, divis)
	}

	for i := range vetor1 {
		div := false

		fmt.Printf("\nNúmero %d (posição %d):\n", vetor1[i], i)

		for k := range vetor2 {
			if vetor1[i]%vetor2[k] == 0 {
				div = true
				fmt.Printf("  -> Divisível por %d na posição %d\n", vetor2[k], k)
			}
		}

		if div == false {
			fmt.Println("  -> Não possui divisores no vetor2")
		}
	}
}
