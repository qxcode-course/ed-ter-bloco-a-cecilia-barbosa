package main

import "fmt"

func main() {
	nConsultas, nBuscas := 0, 0

	fmt.Scan(&nConsultas)
	consulta := ""
	matriz := make(map[string]int)

	for i := 0; i < nConsultas; i++ {
		fmt.Scanln(&consulta)
		matriz[consulta]++
	}

	fmt.Scan(&nBuscas)
	busca := ""
	result := make([]int, nBuscas)

	for i := 0; i < nBuscas; i++ {
		fmt.Scanln(&busca)
		result[i] = matriz[busca]
	}

	saida_f := fmt.Sprintf("%v", result)
	fmt.Println(saida_f[1 : len(saida_f)-1])

}
