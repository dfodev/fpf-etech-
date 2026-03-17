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
	pessoa = append(pessoas, pessoa)

	jsonBytes, _ := json.MarshalIndent(pessoas, "", "")
	os.WriteFile(arquivo, jsonBytes, 0644)
	fmt.Println("adicionado com sucesso!")

}

func buscarPessoa() {
	arquivo := "pessoa.json"
	dados, _ := os.ReadFile(arquivo)

	var pessoas []Pessoa
	json.Unmarshal(dados, &pessoas)

	for i := range pessoas {
		if pessoas[i].Pessoa == "alexandre" {
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
		if pessoas[i].Pessoa == "alexandre" {
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
		if pessoas[i].Pessoa != "alexandre" {
			novaPessoa = append(novaPessoa, pessoas[i])
		}
	}

	jsonBytes, _ := json.MarshalIndent(novaPessoa, "", "")
	os.WriteFile(arquivo, jsonBytes, 0644)
}
func main() {

}
