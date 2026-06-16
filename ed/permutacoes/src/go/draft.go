package main

import "fmt"

func permutacao(letras []rune, usado []bool, nSequencia []rune) []string {
	if len(nSequencia) == len(letras) {
		return []string{string(nSequencia)}
	}

	result := []string{}
	for i := 0; i < len(letras); i++ {
		if usado[i] {
			continue
		}
		usado[i] = true
		result = append(result, permutacao(letras, usado, append(nSequencia, letras[i]))...)
		usado[i] = false
	}

	return result
}

func main() {
	var entrada string
	fmt.Scanln(&entrada)

	letras := []rune(entrada)
	usado := make([]bool, len(letras))

	result := permutacao(letras, usado, []rune{})

	for _, perm := range result {
		fmt.Println(perm)
	}
}
