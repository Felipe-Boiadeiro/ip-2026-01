package main

import "fmt"

func inverter(arr []int) []int {
	inverso := make([]int, len(arr))
	for i := range arr {
		inverso[i] = arr[len(arr)-1-i]
	}
	return inverso
}

func main() {
	n := 0
	fmt.Print("Digite a quantidade de elementos: ")
	fmt.Scan(&n)

	vetor := make([]int, n)

	fmt.Println("Digite os elementos:")
	for i := 0; i < n; i++ {
		fmt.Scan(&vetor[i])
	}

	fmt.Println(inverter(vetor))

}
