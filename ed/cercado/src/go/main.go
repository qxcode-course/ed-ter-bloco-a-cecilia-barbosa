package main

import (
	"bufio"
	"fmt"
	"os"
)

func dfs(grid [][]byte, i int, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != 'O' {
		return
	}

	grid[i][j] = '0' // precisa tornar invalido

	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	l, c := len(board), len(board[0])

	for i := 0; i < l; i++ {
		dfs(board, i, 0)
		dfs(board, i, c-1)
	}

	for i := 0; i < c; i++ {
		dfs(board, 0, i)
		dfs(board, l-1, i)
	}

	for i := 0; i < l; i++ {
		for j := 0; j < c; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '0' {
				board[i][j] = 'O'
			}
		}
	}
}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
