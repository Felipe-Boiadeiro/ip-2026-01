package main

import "fmt"

func main() {

	n1, n2, media, somaMedias := .0, .0, .0, .0
	totalAlunos, reprovados, exame, aprovados := 0, 0, 0, 0
	notas := []float64{}

	fmt.Print("Informe o total de alunos: ")
	fmt.Scan(&totalAlunos)

	for i := 1; i <= totalAlunos; i++ {
		fmt.Print("Informe N1: ")
		fmt.Scan(&n1)
		fmt.Print("Informe N2: ")
		fmt.Scan(&n2)

		media = (n1 + n2) / 2
		somaMedias += media

		notas = append(notas, media)

		fmt.Printf("A média do %vº aluno é %.2f\n", i, media)

		if media <= 3 {
			fmt.Println("Reprovado")
		} else if media >= 7 {
			fmt.Println("Aprovado")
		} else {
			fmt.Println("Exame")
		}
	}

	for i := range notas {
		if notas[i] <= 3 {
			reprovados++
		} else if notas[i] >= 7 {
			aprovados++
		} else {
			exame++
		}
	}

	fmt.Println("Total de aprovados:", aprovados)
	fmt.Println("Total de exame:", exame)
	fmt.Println("Total de reprovados:", reprovados)

	fmt.Printf("Média da classe: %.2f\n", somaMedias/float64(totalAlunos))
}
