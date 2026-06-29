package main

import "fmt"

func sudoku(matriz [][]rune, l int, c int) bool {
	if noLimite(matriz, l, c) == false || matriz[l][c] != '.' {
		return false
	}

	N := len(matriz)

	if l == N-1 && c == N-1 {
		matriz[l][c] = 'X'
		return true
	}

	matriz[l][c] = 'X'

	if sudoku(matriz, l, c+1) {
		return true
	}

	if sudoku(matriz, l+1, c) {
		return true
	}

	if sudoku(matriz, l, c-1) {
		return true
	}

	if sudoku(matriz, l-1, c) {
		return true
	}

	matriz[l][c] = '.'

	return false
}

func noLimite(matriz [][]rune, l, c int) bool {
	N := len(matriz)
	return l >= 0 && l < N && c >= 0 && c < N
}

func podeCaminhar(matriz [][]rune, l, c int) bool {
	return noLimite(matriz, l, c) && matriz[l][c] == '.'
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

	if sudoku(matriz, 0, 0) {
		for _, linha := range matriz {
			fmt.Println(string(linha))
		}
	}
}
