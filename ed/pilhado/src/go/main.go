package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func main() {
	nl, nc := 0, 0
	fmt.Scan(&nl, &nc)

	scanner := bufio.NewScanner(os.Stdin)

	grid := make([][]rune, nl)
	var start, end Pos

	for i := 0; i < nl; i++ {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid[i] = line
		for j, elem := range line {
			if elem == 'I' {
				start = Pos{i, j}
			} else if elem == 'F' {
				end = Pos{i, j}
			}
		}
	}

	visitado := make([][]bool, nl)
	for i := range visitado {
		visitado[i] = make([]bool, nc)
	}

	cam := NewStack[Pos]()
	pass := NewStack[Pos]()
	cam.Push(start)

	for !cam.IsEmpty() {
		atual := cam.Top()
		visitado[atual.l][atual.c] = true

		if atual == end {
			break
		}

		viz := []Pos{
			{atual.l, atual.c - 1},
			{atual.l - 1, atual.c},
			{atual.l, atual.c + 1},
			{atual.l + 1, atual.c},
		}

		var prox *Pos
		for _, v := range viz {
			if v.l >= 0 && v.l < nl && v.c >= 0 && v.c < nc && grid[v.l][v.c] != '#' && !visitado[v.l][v.c] {
				prox = &v
				break
			}
		}

		if prox != nil {
			cam.Push(*prox)
		} else {
			pass.Push(cam.Pop())
		}
	}
	for !cam.IsEmpty() {
		pop := cam.Pop()
		grid[pop.l][pop.c] = '.'

		if grid[pop.l][pop.c] != 'I' && grid[pop.l][pop.c] != 'F' {
			grid[pop.l][pop.c] = '.'
		}
	}

	for _, print := range grid {
		fmt.Println(string(print))
	}
}
