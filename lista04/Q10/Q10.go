package main

import "fmt"

func main() {
	fibonacci := []int{1, 1}

	for i := 2; i < 50; i++ {
		fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
	}

	fmt.Println(fibonacci)
}
