package main

import (
	"bufio"
	"fmt"
	"os"
)

func dfs(grid [][]byte, l, c int) {
	if l < 0 || l >= len(grid) || c < 0 || c >= len(grid[0]) || grid[l][c] == '0' {
		return
	}

	grid[l][c] = '0'

	dfs(grid, l+1, c)
	dfs(grid, l-1, c)
	dfs(grid, l, c+1)
	dfs(grid, l, c-1)
}

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {
	cont := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				cont++
			}
		}
	}
	return cont
}

// Não modifique a função main

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}
