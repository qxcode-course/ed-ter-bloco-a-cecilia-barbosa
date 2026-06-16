package main

import "fmt"

func condicao(tab []int, linha, coluna int) bool {
	for i := 0; i < linha; i++ {
		if tab[i] == coluna { // msm coluna
			return false
		}

		if tab[i]-coluna == i-linha { // msm diagonal
			return false
		}

		if difCol := tab[i] - coluna; difCol == -(i - linha) { // outra diagonal
			return false
		}
	}
	return true
}

func dist(tab []int, linha, n int) int {
	if linha == n {
		return 1
	}

	count := 0
	for col := 0; col < n; col++ {
		if condicao(tab, linha, col) {
			tab[linha] = col
			count += dist(tab, linha+1, n)
			tab[linha] = -1
		}
	}
	return count
}

func main() {
	n := 0
	fmt.Scanln(&n)

	tab := make([]int, n)
	for i := range tab {
		tab[i] = -1
	}

	count := dist(tab, 0, n)
	fmt.Printf("%d\n", count)
}
