package main

import (
	"fmt"
)

func somaAteN(n int) int {
	for n <= 1 {
		return n
	}
	return n + somaAteN(n-1)
}

func main() {
	var num int

	fmt.Print("Inserir numero: ")
	fmt.Scan(&num)

	fmt.Printf("Soma dos numero: (%d) = %d\n", num, somaAteN(num))
}
