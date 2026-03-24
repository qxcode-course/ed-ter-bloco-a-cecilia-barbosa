package main

import "fmt"

func main() {
	qtd_inicial := 0
	fmt.Scan(&qtd_inicial)

	id := make([]int, qtd_inicial)

	for i := 0; i < qtd_inicial; i++ {
		fmt.Scan(&id)
	}

	left := 0
	sairam := make([]int, left)
	fmt.Scan(&left)
	for i := 0; i < left; i++ {
		fmt.Scan(&sairam)
	}

}
