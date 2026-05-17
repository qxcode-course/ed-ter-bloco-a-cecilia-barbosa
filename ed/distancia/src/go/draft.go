package main

import "fmt"

//import "fmt"

func condicao(line []rune, indice int, num int, lim int, prox int) bool {
	for i := indice + 1; i < indice+1+prox; i++ {
		if i < len(line) && line[i] == rune(num)+'0' {
			return false
		}
	}
	for i := indice - prox; i < indice; i++ {
		if i >= 0 && line[i] == rune(num)+'0' {
			return false
		}
	}
	return true
}

func dist(line []rune, indice int, lim int, prox int) bool {
	if len(line) == indice {
		fmt.Println(string(line))
		return true
	}

	if line[indice] != '.' {
		return dist(line, indice+1, lim, prox)
	}

	for i := 0; i <= lim; i++ {
		if condicao(line, indice, i, lim, prox) {
			line[indice] = rune(i) + '0'

			if dist(line, indice+1, lim, prox) {
				return true
			}
		}
	}
	line[indice] = '.'
	return false
}

func main() {
	limite := 0
	entrada := ""
	indice := 0

	fmt.Scanln(&entrada)
	fmt.Scanln(&limite)

	linha := []rune(entrada)

	dist(linha, indice, limite, limite)
}
