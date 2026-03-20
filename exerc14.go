package main

import (
	"fmt"
)

type Carro struct {
	Marca  string
	Modelo string
	Ano    int
}

func main() {
	meuCarro := Carro{
		Marca:  "Mitsubishi",
		Modelo: "L200",
		Ano:    2024,
	}

	ponteiro := &meuCarro

	fmt.Println("acesso direto:")
	fmt.Printf("marca: %s, modelo: %s, ano: %d\n", meuCarro.Marca, meuCarro.Modelo, meuCarro.Ano)

	fmt.Println("\nacesso via ponteiro:")
	fmt.Printf("marca: %s, modelo: %s, ano: %d", ponteiro.Marca, ponteiro.Modelo, ponteiro.Ano)

}
