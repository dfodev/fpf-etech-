package main

import (
	"fmt"
)

// Carro define a estrutura básica do veículo
type Carro struct {
	Marca string
	Ano   int
	KM    float64
}

// map para armazenar os carros (Chave: Renavan)
var frota = make(map[string]Carro)

// (CRUD)

func cadastrar() {
	var renavan string
	fmt.Print("Renavan: ")
	fmt.Scanln(&renavan)

	if _, existe := frota[renavan]; existe {
		fmt.Println("Erro: Veículo já cadastrado.")
		return
	}

	var c Carro
	fmt.Print("Marca: ")
	fmt.Scanln(&c.Marca)
	fmt.Print("Ano: ")
	fmt.Scanln(&c.Ano)
	fmt.Print("KM: ")
	fmt.Scanln(&c.KM)

	frota[renavan] = c
	fmt.Println("Sucesso!")
}

func listarTodos() {
	if len(frota) == 0 {
		fmt.Println("Frota vazia.")
		return
	}
	for r, c := range frota {
		exibir(r, c)
	}
}

func buscar() {
	var r string
	fmt.Print("Renavan: ")
	fmt.Scanln(&r)

	if c, ok := frota[r]; ok {
		exibir(r, c)
	} else {
		fmt.Println("Não encontrado.")
	}
}

func atualizar() {
	var r string
	fmt.Print("Renavan para atualizar: ")
	fmt.Scanln(&r)

	if _, ok := frota[r]; ok {
		var c Carro
		fmt.Print("Nova Marca: ")
		fmt.Scanln(&c.Marca)
		fmt.Print("Novo Ano: ")
		fmt.Scanln(&c.Ano)
		fmt.Print("Nova KM: ")
		fmt.Scanln(&c.KM)

		frota[r] = c
		fmt.Println("Atualizado!")
	} else {
		fmt.Println("Não existe.")
	}
}

func remover() {
	var r string
	fmt.Print("Renavan para remover: ")
	fmt.Scanln(&r)

	if _, ok := frota[r]; ok {
		delete(frota, r)
		fmt.Println("Removido.")
	} else {
		fmt.Println("Não encontrado.")
	}
}

func exibir(renavan string, c Carro) {
	fmt.Printf("ID: %s | Marca: %s | Ano: %d | KM: %.2f\n", renavan, c.Marca, c.Ano, c.KM)
}

func main() {
	var opcao int

	for {
		fmt.Println("\n--- SISTEMA DE FROTA ---")
		fmt.Println("1. Cadastrar")
		fmt.Println("2. Listar Todos")
		fmt.Println("3. Buscar")
		fmt.Println("4. Atualizar")
		fmt.Println("5. Remover")
		fmt.Println("0. Sair")
		fmt.Print("Escolha: ")

		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			cadastrar()
		case 2:
			listarTodos()
		case 3:
			buscar()
		case 4:
			atualizar()
		case 5:
			remover()
		case 0:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção inválida!")
		}
	}
}
