package main

import "fmt"

func main() {

	sequencia := []interface{}{}
	n := 0
	fmt.Scan(&n)

	for {

		if n%16 == 10 {
			sequencia = append(sequencia, "A")
		} else if n%16 == 11 {
			sequencia = append(sequencia, "B")
		} else if n%16 == 12 {
			sequencia = append(sequencia, "C")
		} else if n%16 == 13 {
			sequencia = append(sequencia, "D")
		} else if n%16 == 14 {
			sequencia = append(sequencia, "E")
		} else if n%16 == 15 {
			sequencia = append(sequencia, "F")
		} else {
			sequencia = append(sequencia, n%16)
		}

		n /= 16

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
