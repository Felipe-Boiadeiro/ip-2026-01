package main

import "fmt"

func main() {
	num := 0
	meses := 0

	empregados := []int{}
	tempo := map[int]int{}

	for {
		fmt.Scan(&num, &meses)

		if num == 0 && meses == 0 {
			break
		}

		empregados = append(empregados, num)
		tempo[num] = meses
	}

	m1, m2, m3 := 9999, 9999, 9999
	e1, e2, e3 := 0, 0, 0

	for _, emp := range empregados {
		t := tempo[emp]

		if t < m1 {
			m3, e3 = m2, e2
			m2, e2 = m1, e1
			m1, e1 = t, emp
		} else if t < m2 {
			m3, e3 = m2, e2
			m2, e2 = t, emp
		} else if t < m3 {
			m3, e3 = t, emp
		}
	}

	fmt.Println("Empregados mais recentes:")
	fmt.Printf("1º: %d (%d meses)\n", e1, m1)
	fmt.Printf("2º: %d (%d meses)\n", e2, m2)
	fmt.Printf("3º: %d (%d meses)\n", e3, m3)
}
