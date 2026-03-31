package main

import "fmt"

func main() {
	nConsultas, nBuscas := 0, 0

	fmt.Scan(&nConsultas)
	consulta := make([]string, nConsultas)
	for i := 0; i < nConsultas; i++ {
		fmt.Scanln(&consulta[i])
	}

	fmt.Scan(&nBuscas)
	busca := make([]string, nBuscas)
	for i := 0; i < nBuscas; i++ {
		fmt.Scanln(&busca[i])
	}

	matriz := make(map[string]bool)
}
