package main

import "fmt"

func main() {
	qtd_inicial := 0
	fmt.Scan(&qtd_inicial)

	id := make([]int, qtd_inicial) // identificadores das pessoas
	for i := range id {
		fmt.Scan(&id[i]) // unico tar
	}

	left := 0 // quantas pessoas saíram
	fmt.Scan(&left)

	sairam := make([]int, left) // id das pessoas que saíram
	for i := range sairam {     // range serve para percorrer todo o vetor sem condições ou incremento
		fmt.Scan(&sairam[i])
	}

	idExcluido := make(map[int]bool)     // guardar identificação das que saíram em map
	ordem := make([]int, 0, qtd_inicial) // ordem final

	for _, sairo := range sairam { // identifica todas que sairam e guarda
		idExcluido[sairo] = true
	}

	for _, nova := range id {
		if !idExcluido[nova] { // guarda tudo que for difente do que esta excluido em ordem
			ordem = append(ordem, nova) // adiciona o valor presente em nova e ajusta o tamanho
		}
	}

	saida := fmt.Sprintf("%v", ordem)  // guarda em saída os valores de ordem
	fmt.Print(saida[1 : len(saida)-1]) // len aqui serve pra excluir os ]
	fmt.Println(" ")
}
