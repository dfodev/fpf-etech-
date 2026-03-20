package main

import (
	"fmt"
)

type Carro struct {
	Nome    string
	Proximo *Carro
}

func percorre(lista *Carro, novoNome string) {
	if lista == nil {
		return
	}
	atual := lista

	for atual.Proximo != nil {
		fmt.Println("Carro atual:", atual.Nome)
		atual = atual.Proximo
	}
	fmt.Println("Carro atual:", atual.Nome)

	atual.Proximo = &Carro{Nome: novoNome, Proximo: nil}
	fmt.Println("Carro adicionado:", novoNome)
}

func main() {

	carro3 := &Carro{Nome: "Renault", Proximo: nil}
	carro2 := &Carro{Nome: "Fiat", Proximo: carro3}
	carro1 := &Carro{Nome: "Ford", Proximo: carro2}

	percorre(carro1, "Chevrolet")

}
