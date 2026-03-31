package main

import "fmt"

func main() {
	nConsultas, nBuscas := 0, 0
    
	fmt.Scan(&nConsultas)
	consulta := make([]string, nConsultas)
	matriz := make(map[string]int)
	
	for i := 0; i < nConsultas; i++ {
		fmt.Scan(&consulta[i])
		matriz[consulta[i]]++
	}

	fmt.Scan(&nBuscas)
	busca := make([]string, nBuscas)
	result := make([]int, nBuscas)

	for i := 0; i < nBuscas; i++ {
		fmt.Scan(&busca[i])
		result[i] = matriz[busca[i]]
	}

	saida_f := fmt.Sprintf("%v", result)
	fmt.Println(saida_f[1 : len(saida_f)-1])
}