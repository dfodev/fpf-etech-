package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	var p *int

	fmt.Println("Valor de A:", a)
	fmt.Println("Valor de B:", b)

	b = a
	p = &a

	fmt.Println("Endereço de memória de a (p):", p)
	fmt.Println("Valor armazenado no endereço (*p):", *p)

	fmt.Println("Novo valor de B:", b)

	*p = 30
	fmt.Println("Novo valor de A:", a)

}
