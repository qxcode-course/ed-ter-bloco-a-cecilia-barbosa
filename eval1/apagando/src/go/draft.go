package main

import "fmt"

func main() {
	qtd_pessoas := 0
	fmt.Scan(&qtd_pessoas)

	fila := make([]int, qtd_pessoas)

	for i := 0; i < qtd_pessoas; i++ {
		fmt.Scan(fila[i])
	}
	qtd_sairam := 0
	fmt.Scan(&qtd_sairam)

	sairam := make([]int, qtd_sairam)
	for i := 0; i < qtd_sairam; i++ {
		fmt.Scan(sairam[i])
	}

	temp := make(map[int]bool)
	ordem := make([]int, 0, qtd_pessoas)

	for _, pessoas := range fila {
		temp[pessoas] = true
	}

	for _, nova := range fila {
		if !temp[nova] {
			ordem = append(ordem, nova)
		}
	}

	for i := 0; i < len(ordem); i++ {
		fmt.Printf("%v ", ordem[i])
	}
}
