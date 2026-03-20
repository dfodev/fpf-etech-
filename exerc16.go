package main

import (
	"fmt"
)

type Carro struct {
	Nome    string
	Proximo *Carro
}

func percorre(lista *Carro) {
	atual := lista

	for atual != nil {
		fmt.Println("Carro atual:", atual.Nome)
		atual = atual.Proximo
	}
	atual = lista

}

func main() {

	carro3 := &Carro{Nome: "Renault", Proximo: nil}
	carro2 := &Carro{Nome: "Fiat", Proximo: carro3}
	carro1 := &Carro{Nome: "Ford", Proximo: carro2}

	percorre(carro1)
}
