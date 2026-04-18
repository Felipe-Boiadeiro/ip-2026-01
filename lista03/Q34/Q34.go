package main

import "fmt"

func mdc(x, y int) int {
	for y != 0 {

		x, y = y, x%y
	}
	return x

}

func main() {

	mmc := 0

	n1, n2 := 0, 0

	fmt.Scan(&n1)
	fmt.Scan(&n2)

	mmc = n1 * n2 / mdc(n1, n2)

	fmt.Println(mmc)
}
