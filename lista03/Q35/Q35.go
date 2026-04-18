package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)

	sequencia := []int{}

	for {
		if n%2 == 0 {
			sequencia = append(sequencia, 0)
			n /= 2
		} else {
			sequencia = append(sequencia, 1)
			n /= 2
		}

		if n < 1 {
			break
		}
	}
	i := 0
	x := len(sequencia) - 1

	for i < x {
		sequencia[i], sequencia[x] = sequencia[x], sequencia[i]
		i++
		x--
	}

	fmt.Println(sequencia)
}
