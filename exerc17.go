package main

import (
	"fmt"
)

type Node struct {
	Index             int
	Valor             interface{}
	Anterior, Proximo *Node
}

func (n *Node) Inserir(valor interface{}) {
	atual := n
	contador := 0

	for atual.Proximo != nil {
		atual = atual.Proximo
		contador++
	}

	novoNo := &Node{
		Valor:    valor,
		Index:    contador + 1,
		Anterior: atual,
	}
	atual.Proximo = novoNo
}

func (n *Node) RemoverPorIndex(id int) {
	if id == 0 {
		fmt.Println("Erro: o primeiro NÓ(head), não pode ser removido.")
		return
	}

	atual := n
	for atual != nil {
		if atual.Index == id {
			if atual.Anterior != nil {
				atual.Anterior.Proximo = atual.Proximo
			}
			if atual.Proximo != nil {
				atual.Proximo.Anterior = atual.Anterior
			}
			n.reindexar()
			fmt.Printf("Nó removido! ID: %d\n", id)
			return
		}
		atual = atual.Proximo
	}
}

func (n *Node) reindexar() {
	atual := n
	i := 0
	for atual != nil {
		atual.Index = i
		i++
		atual = atual.Proximo
	}
}

func (n *Node) Listar() {
	atual := n
	for atual != nil {
		fmt.Printf("ID: %d, Valor: %v\n", atual.Index, atual.Valor)
		atual = atual.Proximo
	}
	fmt.Println("---------------------")

}

func main() {
	head := &Node{Index: 0, Valor: "Início (Head)"}

	head.Inserir("Tarefa 1")
	head.Inserir("Tarefa 2")
	head.Inserir("Tarefa 3")
	head.Inserir("Tarefa 4")
	head.Inserir("Tarefa 5")
	head.Inserir("Tarefa 6")
	head.Inserir("Tarefa 7")
	head.Inserir("Tarefa 8")

	fmt.Print("Lista original: ")
	head.Listar()

	head.RemoverPorIndex(6)

	fmt.Print("Lista após remoção: ")
	head.Listar()

}
