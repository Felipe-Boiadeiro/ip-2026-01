package main

import "fmt"

var hor, min, seg, tf int

func main() {

	fmt.Println("Informe a hora: ")

	fmt.Scan(&hor)

	fmt.Println("Informe os minutos: ")

	fmt.Scan(&min)

	fmt.Println("Informe os segundos: ")

	fmt.Scan(&seg)

	tf = hor*3600 + min*60 + seg

	fmt.Printf("O tempo em segundos é = %d", tf)

}
