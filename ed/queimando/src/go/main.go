package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func burnTrees(grid [][]rune, l, c int) {
	if l < 0 || l >= len(grid) || c < 0 || c >= len(grid[0]) || grid[l][c] != '#' {
		return
	}
	stack := NewStack[Pos]()
	stack.Push(Pos{l, c})

	for !stack.IsEmpty() {
		elem := stack.Pop()
		lin, col := elem.l, elem.c

		if grid[lin][col] == '#' {
			grid[lin][col] = 'o'

			vizinhos := []Pos{
				{lin - 1, col},
				{lin + 1, col},
				{lin, col - 1},
				{lin, col + 1},
			}

			for _, v := range vizinhos {
				if v.l >= 0 && v.l < len(grid) && v.c >= 0 && v.c < len(grid[0]) && grid[v.l][v.c] == '#' {
					stack.Push(v)
				}
			}
		}
	}
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

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
