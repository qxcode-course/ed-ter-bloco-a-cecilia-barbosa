package main

import "fmt"

func main() {
	entrada := ""
	fmt.Scanln(&entrada)

	texto := []rune{}
	posicao := 0
	// REFAZENDO
	for _, c := range entrada {
		switch c {
		case '>':
			if posicao < len(texto) {
				posicao++
			}
		case '<':
			if posicao > 0 {
				posicao--
			}
		case 'B':
			if posicao > 0 {
				var temp []rune
				for i := 0; i < posicao-1; i++ {
					temp = append(temp, texto[i])
				}

				for i := posicao; i < len(texto); i++ {
					temp = append(temp, texto[i])
				}

				texto = temp
				posicao--
			}
		case 'D':
			if posicao < len(texto) {
				var temp []rune
				for i := 0; i < posicao; i++ {
					temp = append(temp, texto[i])
				}
				//pula
				for i := posicao + 1; i < len(texto); i++ {
					temp = append(temp, texto[i])
				}
				texto = temp
			}
		case 'R':
			var temp []rune
			for i := 0; i < posicao; i++ {
				temp = append(temp, texto[i])
			}

            temp = append(temp, '\n')

			for i := posicao; i < len(texto); i++ {
				temp = append(temp, texto[i])
			}

			texto = temp
			posicao++
		default:
			var temp []rune
			for i := 0; i < posicao; i++ {
				temp = append(temp, texto[i])
			}

			temp = append(temp, c)
			for i := posicao; i < len(texto); i++ {
				temp = append(temp, texto[i])
			}
			texto = temp
			posicao++
		}
	}

	for i, letra := range texto {
		if i == posicao {
			fmt.Print("|")
		}
		fmt.Printf("%c", letra)
	}

	if posicao == len(texto) {
		fmt.Print("|\n")
	} else {
		fmt.Println()
	}
}
