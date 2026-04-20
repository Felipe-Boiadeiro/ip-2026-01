package main

import "fmt"

func main() {
	cpf := ""
	fmt.Scan(&cpf)

	digitos := []int{}

	if len(cpf) > 9 {
		fmt.Println("Inválido")
		return
	}

	for _, v := range cpf {
		if v >= '0' && v <= '9' {
			digitos = append(digitos, int(v-'0'))
		}
	}

	soma := 0
	peso := 10

	for i := 0; i < 9; i++ {
		soma += digitos[i] * peso
		peso--
	}

	if soma%11 < 2 {
		digitos = append(digitos, 0)
	} else {
		digitos = append(digitos, 11-soma%11)
	}

	soma2 := 0
	peso2 := 11
	for i := 0; i < 10; i++ {
		soma2 += digitos[i] * peso2
		peso2--
	}

	if soma2%11 < 2 {
		digitos = append(digitos, 0)
	} else {
		digitos = append(digitos, 11-soma2%11)
	}

	fmt.Println(digitos)
}
