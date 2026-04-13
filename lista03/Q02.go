package main

import "fmt"

func main() {
	var soma, media int
	for x := 50; x <= 70; x += 2 {
		soma += x
	}
	media = soma / 11
	fmt.Println(media)
}
