package main

import (
	"bufio"
	"fmt"
	"os"
)

func burnTrees(grid [][]rune, l, c int) {
	numL := len(grid)
	numC := len(grid[0])

	if c < 0 || c >= numC || l < 0 || l >= numL { // se estiver fora da matriz, retorne
		return
	}
	if grid[l][c] != '#' { // se o elemento atual não for uma arvore, retorne
		return
	}

	grid[l][c] = 'o'        // queime a arvore colocando o caractere 'o' na posição atual
	burnTrees(grid, l, c+1) // chame a recursão para todos os 4 vizinhos
	burnTrees(grid, l, c-1)
	burnTrees(grid, l+1, c)
	burnTrees(grid, l-1, c)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(grid [][]rune) {
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
