package main

import (
	"fmt"
	"strconv"
)

func binario(x int) string {
	bin := strconv.FormatInt(int64(x), 2)
	return bin

}

func main() {
	n := 0
	fmt.Scan(&n)
	fmt.Println(binario(n))

}
