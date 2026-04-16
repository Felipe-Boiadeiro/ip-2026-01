package main

import "fmt"

func main() {
	numeros := []int{}
	n, soma, contador, media := 0, 0, 0, .0

	for {
		fmt.Print("Digite um número (caso queira parar digite 30000): ")
		fmt.Scan(&n)
		if n == 30000 {
			break
		}
		numeros = append(numeros, n)

	}
	for i := range numeros {
		soma += numeros[i]
	}
	for i := 0; i < len(numeros); i++ {
		for x := 0; x < len(numeros)-1; x++ {
			if numeros[x] < numeros[x+1] {
				numeros[x], numeros[x+1] = numeros[x+1], numeros[x]
			}
		}
	}
	for i := range numeros {
		if numeros[i]%2 == 0 {
			media += float64(numeros[i])
			contador++
		}
	}
	fmt.Println(soma)
	fmt.Println(len(numeros))
	fmt.Printf("%.2f\n", float64(soma)/float64(len(numeros)))
	fmt.Println(numeros[0])
	fmt.Println(numeros[len(numeros)-1])
	fmt.Printf("%.2f\n", media/float64(contador))
	fmt.Printf("%.2f", float64(len(numeros)-contador)/float64(len(numeros)))
}
