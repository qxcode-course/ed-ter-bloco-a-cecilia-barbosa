package main

import "fmt"

func balanceada(entrada string) string {
	pilha := []rune{}
	for _, temp := range entrada {
		if temp == '(' || temp == '[' {
			pilha = append(pilha, temp)
		} else if temp == ')' || temp == ']' {
			if len(pilha) == 0 {
				return "nao balanceado"
			}
			cima := pilha[len(pilha)-1]
			pilha = pilha[:len(pilha)-1]

			if (temp == ')' && cima != '(') || (temp == ']' && cima != '[') {
				return "nao balanceado"
			}
		}
	}
	if len(pilha) == 0 {
		return "balanceado"
	}
	return "nao balanceado"
}

func main() {
	entrada := ""

	fmt.Scan(&entrada)
	fmt.Println(balanceada(entrada))
}
