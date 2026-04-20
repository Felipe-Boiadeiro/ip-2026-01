package main

import "fmt"

func main() {
	var idGordo, idMagro int
	var pesoGordo, pesoMagro float64

	for i := 0; i < 90; i++ {
		var id int
		var peso float64
		fmt.Scan(&id, &peso)

		if i == 0 {
			idGordo, pesoGordo = id, peso
			idMagro, pesoMagro = id, peso
		} else {
			if peso > pesoGordo {
				idGordo, pesoGordo = id, peso
			}
			if peso < pesoMagro {
				idMagro, pesoMagro = id, peso
			}
		}
	}

	fmt.Println("Mais gordo: ID", idGordo, "Peso:", pesoGordo)
	fmt.Println("Mais magro: ID", idMagro, "Peso:", pesoMagro)
}
