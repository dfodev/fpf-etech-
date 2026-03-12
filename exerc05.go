package main

import (
	"fmt"
	"os"
	"time"
)

func log(resultado string) {
	file, err := os.OpenFile("log_Operacoes.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Erro ao abrir log:", err)
		return
	}
	defer file.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	linha := fmt.Sprintf("[%s]: %s\n", timestamp, resultado)

	file.WriteString(linha)
}

func fila() {
	var fila []int

	fmt.Printf("insira 5 elementos\n")

	for i := 1; i < 5; i++ {
		var valor int
		fmt.Printf("Digite o elemento: ")
		fmt.Scanf("%d", &valor)

		fila = append(fila, valor)
		log(fmt.Sprintf("Fila inserida: "))
	}
	fmt.Println("Fila final:", fila)

	for len(fila) > 0 {
		removido := fila[0]
		fila = fila[1:]

		msg := fmt.Sprint("fila - removido: ", removido)
		fmt.Println(msg)
		log(msg)
	}
}

func pilha() {
	var pilha []int

	fmt.Printf("insira 5 elementos\n")

	for i := 1; i < 5; i++ {
		var valor int
		fmt.Printf("Elemento para o topo: ")
		fmt.Scanf("%d", &valor)
		// Na pilha, adicionamos ao final, mas a saída lógica é inversa
		pilha = append(pilha, valor)
	}
	fmt.Print("Desempilhando (LIFO): ", pilha)

	fmt.Println("\nRemovendo elementos da PILHA:")
	for len(pilha) > 0 {
		indiceUltimo := len(pilha) - 1
		removido := pilha[indiceUltimo] // Captura o último elemento
		pilha = pilha[:indiceUltimo]    // Remove o último elemento da fatia

		msg := fmt.Sprintf("PILHA - Remoção: %d", removido)
		fmt.Println(msg)
		log(msg)
	}
}

func main() {
	var opcao int

	for {
		fmt.Println("\n--- MENU DE MANIPULAÇÃO ---")
		fmt.Println("1) FILA")
		fmt.Println("2) PILHA")
		fmt.Println("-1) SAIR")
		fmt.Print("Digite a opção: ")
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			fila()
			fmt.Println("Manipulada fila")
			log("fila:")
		case 2:
			pilha()
			fmt.Println("Manipulada pilha")
			log("pilha:")

		case 3:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção inválida!")
		}
	}
}
