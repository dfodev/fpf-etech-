package main

import (
	"fmt"
)

func matrizes(matriz [][]int) {
	for i := range matriz {
		for j := range matriz[i] {
			if i == j {
				matriz[i][j] = 1
			} else {
				matriz[i][j] = 0
			}
		}
	}
}

func main() {
	minhaMatriz := [][]int{
		{1, 3, 10},
		{0, 1, 5},
		{3, 0, 1},
	}
	matrizes(minhaMatriz)

	fmt.Println("Matriz Identidade resultante:")
	for _, linha := range minhaMatriz {
		fmt.Println(linha)
	}
}
