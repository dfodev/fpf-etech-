package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func registrarLog(acao string, resultado string) {
	file, err := os.OpenFile("log_sistema.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Erro ao abrir log:", err)
		return
	}
	defer file.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	linha := fmt.Sprintf("[%s] [%s] [%s]: \n", timestamp, acao, resultado)
	file.WriteString(linha)

}

func mostrarElementos(lista []int) {
	fmt.Printf("\nEstrutura atual: %v\n", lista)
}

func manipularFila(dados *[]int) {
	registrarLog("Menu", "Acesso SubMenu Fila: ")
	var opcao int

	for {
		fmt.Println("\n-------SubMenu-------")
		fmt.Println("1) Inserir elementos: ")
		fmt.Println("2) Remover elementos: ")
		fmt.Println("3) Mostrar Topo(Cabeça): ")
		fmt.Println("4) Mostrar Final(Rabo): ")
		fmt.Println("5) Mostrar Fila completa: ")
		fmt.Println("-1) Voltar ao menu principal: ")
		fmt.Println("Escolha: ")
		fmt.Scanln("%d", &opcao)

		if opcao == -1 {
			registrarLog("MENU", "Voltou ao principal vindo de FILA")
			return
		}

		switch opcao {
		case 1:
			valor := rand.Intn(100)
			*dados = append(*dados, valor) // Insere no fim
			registrarLog("FILA", fmt.Sprintf("INSERIR valor=%d | OK", valor))
		case 2:
			if len(*dados) > 0 {
				removido := (*dados)[0]
				*dados = (*dados)[1:] // Remove o primeiro (cabeça)
				registrarLog("FILA", fmt.Sprintf("REMOVER valor=%d | OK", removido))
			} else {
				fmt.Println("Fila vazia!")
			}
		case 3:
			if len(*dados) > 0 {
				fmt.Printf("CABEÇA: %d\n", (*dados)[0])
			} else {
				fmt.Println("Fila vazia!")
			}
		case 4:
			if len(*dados) > 0 {
				fmt.Printf("RABO: %d\n", (*dados)[len(*dados)-1])
			} else {
				fmt.Println("Fila vazia!")
			}
		case 5:
			// Apenas segue para mostrarElementos()
		default:
			fmt.Println("Opção inválida!")
		}
		mostrarElementos(*dados)
	}
}

func manipularPilha(dados *[]int) {
	registrarLog("MENU", "Acessou SUBMENU PILHA")
	var opcao int

	for {
		fmt.Println("\n--- SUBMENU PILHA (LIFO) ---")
		fmt.Println("1) Inserir elemento (Push Aleatório)")
		fmt.Println("2) Remover elemento (Pop)")
		fmt.Println("3) Mostrar Topo")
		fmt.Println("4) Mostrar Base")
		fmt.Println("5) Mostrar Pilha Completa")
		fmt.Println("-1) Voltar ao menu principal")
		fmt.Print("Escolha: ")
		fmt.Scan(&opcao)

		if opcao == -1 {
			registrarLog("MENU", "Voltou ao Principal vindo de PILHA")
			return
		}

		switch opcao {
		case 1:
			valor := rand.Intn(100)
			*dados = append(*dados, valor) // Push (insere no fim)
			registrarLog("PILHA", fmt.Sprintf("PUSH valor=%d | OK", valor))
		case 2:
			if len(*dados) > 0 {
				ultimoIdx := len(*dados) - 1
				removido := (*dados)[ultimoIdx]
				*dados = (*dados)[:ultimoIdx] // Pop (remove o último)
				registrarLog("PILHA", fmt.Sprintf("POP valor=%d | OK", removido))
			} else {
				fmt.Println("Pilha vazia!")
			}
		case 3:
			if len(*dados) > 0 {
				fmt.Printf("TOPO: %d\n", (*dados)[len(*dados)-1])
			} else {
				fmt.Println("Pilha vazia!")
			}
		case 4:
			if len(*dados) > 0 {
				fmt.Printf("BASE: %d\n", (*dados)[0])
			} else {
				fmt.Println("Pilha vazia!")
			}
		case 5:
		default:
			fmt.Println("Opção inválida!")
		}
		mostrarElementos(*dados)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var listaCompartilhada []int

	for {
		var opcao int

		fmt.Println("MENU DE OPÇÕES")
		fmt.Println("1) MANIPULAR LISTA")
		fmt.Println("2) MANIPULAR PILHA")
		fmt.Println("-1) SAIR")
		fmt.Print("ESolha uma estrutura: ")
		fmt.Scanf("%d", &opcao)

		if opcao == -1 {
			registrarLog("SISTEMA", "Programa encerrado pelo usuário")
			fmt.Println("Encerrando sistema...")
			break
		}
		switch opcao {
		case 1:
			manipularFila(&listaCompartilhada)
		case 2:
			manipularPilha(&listaCompartilhada)
		default:
			fmt.Println("Opção inválida! Tente novamente.")
		}
	}
}
