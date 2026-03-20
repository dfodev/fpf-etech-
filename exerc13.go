package main

import (
	"fmt"
)

func dobrar(numero *int) {
	*numero = *numero * 2
}

func main() {
	var numero int = 10

	fmt.Println("Valor inicial:", numero)
	fmt.Println("Endereço:", &numero)

	dobrar(&numero)

	fmt.Println("Valor dobrado:", numero)
	fmt.Println("Endereço:", &numero)

}
