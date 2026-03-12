package main

import (
	"fmt"
)

func quantEntrada(n int) int {
	for n < 10 {
		return 1
	}
	return 1 + quantEntrada(n/10)

}

func main() {
	numero := 53827

	resultado := quantEntrada(numero)

	fmt.Printf("entrada: %d\n", numero)
	fmt.Printf("saida:  %d\n", resultado)
}
