package main

import "fmt"

func main() {
	var num int
	fmt.Print("Imforme um número: ")
	fmt.Scan((&num))

	fat := fatorial(num)

	fmt.Printf("O fatorial de %d é %d\n", num, fat)
}
func fatorial(x int) int {
	if x == 0 {
		return 1
	} else {
		return x * fatorial(x-1)
	}

}
