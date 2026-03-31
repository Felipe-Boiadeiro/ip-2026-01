package main

import "fmt"

func main() {
	var mediaExer, mediaFinal, n1, n2, n3 float64
	var ident int

	fmt.Print("Informe o seu número de identificação: ")
	fmt.Scan(&ident)

	fmt.Print("Informe as notas das verificações 1, 2 e 3 respectivamente: ")
	fmt.Scan(&n1, &n2, &n3)

	fmt.Print("Informe a média dos exercícios: ")
	fmt.Scan(&mediaExer)

	mediaFinal = (n1 + n2*2 + n3*3 + mediaExer) / 7

	switch {
	case mediaFinal >= 9 && mediaFinal <= 10:
		fmt.Printf("Aluno: %v\n Nota 1= %v\n Nota 2= %v\n Nota 3= %v\n Média dos execícios: %.2f\n Média de aproveitamento: %.2f\n Conceito Final: A\n APROVADO\n", ident, n1, n2, n3, mediaExer, mediaFinal)

	case mediaFinal >= 7.5 && mediaFinal < 9:
		fmt.Printf("Aluno: %v\n Nota 1= %v\n Nota 2= %v\n Nota 3= %v\n Média dos execícios: %.2f\n Média de aproveitamento: %.2f\n Conceito Final: B\n APROVADO\n", ident, n1, n2, n3, mediaExer, mediaFinal)

	case mediaFinal >= 6 && mediaFinal < 7.5:
		fmt.Printf("Aluno: %v\n Nota 1= %v\n Nota 2= %v\n Nota 3= %v\n Média dos execícios: %.2f\n Média de aproveitamento: %.2f\n Conceito Final: C\n APROVADO\n", ident, n1, n2, n3, mediaExer, mediaFinal)

	case mediaFinal >= 4 && mediaFinal < 6:
		fmt.Printf("Aluno: %v\n Nota 1= %v\n Nota 2= %v\n Nota 3= %v\n Média dos execícios: %.2f\n Média de aproveitamento: %.2f\n Conceito Final: D\n REPROVADO\n", ident, n1, n2, n3, mediaExer, mediaFinal)

	case mediaFinal < 4:
		fmt.Printf("Aluno: %v\n Nota 1= %v\n Nota 2= %v\n Nota 3= %v\n Média dos execícios: %.2f\n Média de aproveitamento: %.2f\n Conceito Final: E\n REPROVADO\n", ident, n1, n2, n3, mediaExer, mediaFinal)
	}
}
