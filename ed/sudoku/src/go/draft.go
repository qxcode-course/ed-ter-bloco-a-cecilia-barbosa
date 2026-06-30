package main

import "fmt"

func sudoku(matriz [][]rune, indice int) bool {
	N := len(matriz)

	if indice == N*N {
		return true
	}

	l := indice / N
	c := indice % N

	if matriz[l][c] != '.' {
		return sudoku(matriz, indice+1)
	}

	for v := '1'; v <= rune('0'+N); v++ {
		if podeColocar(matriz, l, c, v) {
			matriz[l][c] = v
			if sudoku(matriz, indice+1) {
				return true
			}

			matriz[l][c] = '.'
		}
	}

	return false
}

func podeColocar(matriz [][]rune, l, c int, v rune) bool {
	return !linha(matriz, l, v) && !coluna(matriz, c, v) && !quadro(matriz, l, c, v)
}

func linha(matriz [][]rune, l int, v rune) bool {
	for _, x := range matriz[l] {
		if x == v {
			return true
		}
	}
	return false
}

func coluna(matriz [][]rune, c int, v rune) bool {
	for _, linha := range matriz {
		if linha[c] == v {
			return true
		}
	}
	return false
}

func quadro(matriz [][]rune, l, c int, v rune) bool {
	N := len(matriz)
	tam := 2
	if N == 9 {
		tam = 3
	}

	lz := (l / tam) * tam
	cz := (c / tam) * tam

	for i := 0; i < tam; i++ {
		for j := 0; j < tam; j++ {
			if matriz[lz+i][cz+j] == v {
				return true
			}
		}
	}
	return false
}

func main() {
	N := 0
	fmt.Scan(&N)

	matriz := make([][]rune, N)
	for i := range matriz {
		linha := ""
		fmt.Scan(&linha)
		matriz[i] = []rune(linha)
	}

	if sudoku(matriz, 0) {
		for _, linha := range matriz {
			fmt.Println(string(linha))
		}
	}
}
