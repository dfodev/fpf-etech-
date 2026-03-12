package main

import (
	"fmt"
)

func matriz(matriz [][]int) int {
	var acuml = 1
	validador := false

	for _, linha := range matriz {
		for _, valor := range linha {
			// Regra: desconsiderar 0 e negativos
			if valor > 0 {
				acuml *= valor
				validador = true
			}
		}
	}
	if !validador {
		return 0
	}
	return acuml
}

func main() {
	minhaMatriz := [][]int{
		{2, 3, -1},
		{0, 4, 5},
		{-10, 2, 0},
	}
	resultado := matriz(minhaMatriz)
	fmt.Printf("O produto dos elementos (ignorando <= 0) é: %d\n", resultado)

}
