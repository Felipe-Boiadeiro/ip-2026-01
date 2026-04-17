package main

import "fmt"

func main() {

	dividendo := []int{1}
	divisor := []int{1}

	for i := 1; i <= 49; i++ {
		dividendo = append(dividendo, dividendo[i-1]+2)
		divisor = append(divisor, i+1)
	}

	valor := 0.0

	for i := range divisor {
		valor += float64(dividendo[i]) / float64(divisor[i])
	}

	fmt.Println(valor)
}
