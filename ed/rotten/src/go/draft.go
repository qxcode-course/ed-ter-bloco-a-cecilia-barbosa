package main

import (
	"fmt"
)

func bfs() {}

func orangesRotting(grid [][]int) int {
	l, c := len(grid), len(grid[0])
	queue := [][]int{}
	frescas := 0

	for lin := 0; lin < l; lin++ {
		for col := 0; col < c; col++ {
			if grid[lin][col] == 1 {
				frescas++
			}
			if grid[lin][col] == 2 {
				queue = append(queue, []int{lin, col})
			}
		}
	}

	direcoes := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	tempo := 0

	for len(queue) > 0 && frescas > 0 {
		tam := len(queue)

		for i := 0; i < tam; i++ {
			atual := queue[0]
			queue = queue[1:]
			lin, col := atual[0], atual[1]
			for _, dir := range direcoes {
				ln, cl := lin+dir[0], col+dir[1]
				if ln < 0 || ln == l || cl < 0 || cl == c || grid[ln][cl] != 1 {
					continue
				}
				grid[ln][cl] = 2
				queue = append(queue, []int{ln, cl})
				frescas--
			}
		}
		tempo++
	}

	if frescas == 0 {
		return tempo
	}
	return -1
}

func main() {
	l, c := 0, 0
	fmt.Scan(&l, &c)

	grid := make([][]int, l)
	for i := 0; i < l; i++ {
		grid[i] = make([]int, c)
		for j := 0; j < c; j++ {
			fmt.Scan(&grid[i][j])
		}
	}

	resultado := orangesRotting(grid)
	fmt.Println(resultado)
}
