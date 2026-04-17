package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			if i > j {
				fmt.Printf("|%2d,%2d| ", i, j)
			} else {
				fmt.Printf("        ")
			}
		}
		fmt.Println()
	}
}
