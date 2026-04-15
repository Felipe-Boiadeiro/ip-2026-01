package main

import "fmt"

func main() {
	soma, media, contador := 0, .0, 0
	for x := 50; x <= 70; x += 2 {
		soma += x
		contador++
	}
	media = float64(soma) / float64(contador)

	fmt.Printf("A soma é: %v\n A média é: %f", soma, media)

}
