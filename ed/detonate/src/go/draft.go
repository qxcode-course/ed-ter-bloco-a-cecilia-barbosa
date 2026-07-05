package main

import "fmt"

func bfs(graph [][]int, i int, visitado []bool) int {
	visitado[i] = true
	cont := 1
	for _, j := range graph[i] {
		if !visitado[j] {
			cont += bfs(graph, j, visitado)
		}
	}
	return cont
}

func detonate(bombas [][]int) int {
	n := len(bombas)
	graph := make([][]int, n)

	for i := 0; i < n; i++ {
		xi, yi, ri := bombas[i][0], bombas[i][1], bombas[i][2]
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			x2, y2 := bombas[j][0], bombas[j][1]
			dx := xi - x2
			dy := yi - y2
			if dx*dx+dy*dy <= ri*ri {
				graph[i] = append(graph[i], j)
			}
		}
	}
	result := 0
	for i := 0; i < n; i++ {
		visitado := make([]bool, n)
		cont := bfs(graph, i, visitado)
		if cont > result {
			result = cont
		}
	}
	return result
}

func main() {
	n, m := 0, 0

	fmt.Scan(&n, &m)

	bombas := make([][]int, n)

	for i := 0; i < n; i++ {
		var xi, yi, ri int
		fmt.Scan(&xi, &yi, &ri)
		bombas[i] = []int{xi, yi, ri}
	}

	fmt.Println(detonate(bombas))
}
