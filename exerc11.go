package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Pessoa struct {
	Nome           string
	CPF            string
	dataNascimento string
	Idade          int
}

func adicionarPessoa() {
	pessoa := Pessoa{
		Nome:           "daniel",
		CPF:            "123.456.789-00",
		dataNascimento: "1990-01-01",
		Idade:          25,
	}

	arquivo := "pessoa.json"
	dados, _ := os.ReadFile(arquivo)
	fmt.Println(dados)

	var pessoas []Pessoa
	json.Unmarshal(dados, &pessoas)

	fmt.Println(pessoas)
	pessoas = append(pessoas, pessoa)

	jsonBytes, _ := json.MarshalIndent(pessoas, "", "")
	os.WriteFile(arquivo, jsonBytes, 0644)
	fmt.Println("adicionado com sucesso!")

}

func buscarPessoa() {
	arquivo := "pessoa.json"
	var find string
	dados, _ := os.ReadFile(arquivo)
	fmt.Println("Digite o nome da pessoa:")
	fmt.Scanln(&find)
	var pessoas []Pessoa
	json.Unmarshal(dados, &pessoas)

	for i := range pessoas {
		if pessoas[i].Nome == find {
			fmt.Println("CPF ENCONTRADO: ", pessoas[i].CPF)
		}
	}

	jsonBytes, _ := json.MarshalIndent(pessoas, "", "")
	os.WriteFile(arquivo, jsonBytes, 0644)
}

func atualizarPessoa() {
	arquivo := "pessoa.json"
	dados, _ := os.ReadFile(arquivo)

	var pessoas []Pessoa
	json.Unmarshal(dados, &pessoas)

	for i := range pessoas {
		if pessoas[i].Nome == "alexandre" {
			pessoas[i].CPF = " "
		}
	}

	jsonBytes, _ := json.MarshalIndent(pessoas, "", "")
	os.WriteFile(arquivo, jsonBytes, 0644)
}

func deletarPessoa() {
	arquivo := "pessoa.json"
	dados, _ := os.ReadFile(arquivo)

	var pessoas []Pessoa
	json.Unmarshal(dados, &pessoas)

	var novaPessoa []Pessoa
	for i := range pessoas {
		if pessoas[i].Nome != "alexandre" {
			novaPessoa = append(novaPessoa, pessoas[i])
		}
	}

	jsonBytes, _ := json.MarshalIndent(novaPessoa, "", "")
	os.WriteFile(arquivo, jsonBytes, 0644)
}
func main() {
	var opcao int
	for {
		fmt.Println("\n--- Gerenciador de Pessoas ---")
		fmt.Println("1. Adicionar Pessoa")
		fmt.Println("2. Buscar Pessoa")
		fmt.Println("3. Atualizar Pessoa")
		fmt.Println("4. Deletar Pessoa")
		fmt.Println("0. Sair")
		fmt.Print("Escolha uma opção: ")

		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			adicionarPessoa()
		case 2:
			buscarPessoa()
		case 3:
			atualizarPessoa()
		case 4:
			deletarPessoa()
		case 0:
			fmt.Println("Saindo do programa...")
			return // Encerra o programa
		default:
			fmt.Println("Opção inválida! Tente novamente.")

		}
	}
}
