package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Chamando a função GeraSliceSemRepeticao para gerar nosso slice com N numeros
	sliceNumerosAleatorios := GeraSliceSemRepeticao(10000)
	fmt.Println("Tamanho sliceNumerosAleatorios:", len(sliceNumerosAleatorios))

	// Declaração Slices para cada função
	numerosParaBubble := make([]int, len(sliceNumerosAleatorios))
	numerosParaEscolhida := make([]int, len(sliceNumerosAleatorios))

	// Atribuição de valores. Quando estudarmos alocação dinâmica,
	// entenderemos melhor porque usar o copy.
	copy(numerosParaBubble, sliceNumerosAleatorios)
	copy(numerosParaEscolhida, sliceNumerosAleatorios)

	// Temporizador para Bubble Sort
	inicioBubble := time.Now()
	ordenacaoBubble(numerosParaBubble)
	duracaoBubble := time.Since(inicioBubble)

	// Temporizador para ordenacaoQuick (Ordenação Escolhida)
	inicioEscolhida := time.Now()
	ordenacaoQuick(numerosParaEscolhida)
	duracaoEscolhida := time.Since(inicioEscolhida)

	// Printar durações
	fmt.Printf("Tempo ordenacaoBubble: %v\n", duracaoBubble)
	fmt.Printf("Tempo ordenacaoQuick: %v\n", duracaoEscolhida)

	// Comparação de tempos
	if duracaoBubble < duracaoEscolhida {
		fmt.Println("ordenacaoBubble foi mais rápida")
	} else if duracaoEscolhida < duracaoBubble {
		fmt.Println("ordenacaoQuick foi mais rápida")
	} else {
		fmt.Println("Ambas as ordenações tiveram o mesmo tempo")
	}
}

// Gera um slice com tamanho números inteiros únicos de 1 a tamanho, embaralhados.
func GeraSliceSemRepeticao(tamanho int) []int {
	if tamanho <= 0 {
		panic("tamanho deve ser maior que zero")
	}

	rand.Seed(time.Now().UnixNano())
	numeros := make([]int, tamanho)
	for i := 0; i < tamanho; i++ {
		numeros[i] = i + 1
	}
	for i := tamanho - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		numeros[i], numeros[j] = numeros[j], numeros[i]
	}
	return numeros
}

// Implementação do Bubble Sort
func ordenacaoBubble(slice []int) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		// A cada iteração i, o maior elemento "flutua" para o final,
		// por isso o n-i-1 para evitar comparações desnecessárias.
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				// Troca os elementos de posição caso o atual seja maior que o próximo
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

// Função principal da ordenação escolhida (Quick Sort)
func ordenacaoQuick(slice []int) {
	quicksortRecursive(slice, 0, len(slice)-1)
}

// Lógica recursiva do Quick Sort (Dividir e Conquistar)
func quicksortRecursive(slice []int, inicio, fim int) {
	if inicio < fim {
		// Organiza o slice em torno de um pivô e retorna o índice dele
		pivoIndex := particionamento(slice, inicio, fim)

		// Ordena recursivamente os elementos menores (esquerda) e maiores (direita)
		quicksortRecursive(slice, inicio, pivoIndex-1)
		quicksortRecursive(slice, pivoIndex+1, fim)
	}
}

// O particionamento é o motor do Quick Sort: ele organiza os elementos
// sem criar novos slices, apenas trocando-os de lugar (in-place).
func particionamento(slice []int, inicio, fim int) int {
	pivo := slice[fim] // Escolhemos o último elemento como pivô
	i := inicio - 1

	for j := inicio; j < fim; j++ {
		// Se o elemento atual for menor ou igual ao pivô, movemos para a esquerda
		if slice[j] <= pivo {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}
	// Coloca o pivô na sua posição final correta
	slice[i+1], slice[fim] = slice[fim], slice[i+1]
	return i + 1
}
