package main

import (
	"fmt"
	"sort"
)

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

func ordem(r rune) int {
	if r >= '0' && r <= '9' {
		return 0
	}
	if r >= 'A' && r <= 'Z' {
		return 1
	}
	return 2
}

func main() {
	var entrada string
	fmt.Scanln(&entrada)

	letras := []rune(entrada)
	sort.Slice(letras, func(i, j int) bool {
		a, b := letras[i], letras[j]
		ta := ordem(a)
		tb := ordem(b)
		if ta != tb {
			return ta < tb
		}
		return a < b
	})

	usado := make([]bool, len(letras))
	result := permutacao(letras, usado, []rune{})

	for _, perm := range result {
		fmt.Println(perm)
	}
}
