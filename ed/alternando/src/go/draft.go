package main

import "fmt"

func print_jog(jog []int, espada int) {
	fmt.Print("[ ")
	for i, elem := range jog {
		if elem == 0 {
			continue
		}

		if i == espada {
			if elem > 0 {
				fmt.Printf("%d>", elem)
			} else {
				fmt.Printf("<%d", elem)
			}
		} else {
			fmt.Printf("%d", elem)
		}
		fmt.Print(" ")
	}
	fmt.Printf("]\n")
}

func procura_vivo(jogadores []int, pos int, direcao int) int {
	n := len(jogadores)
	for i := 1; i <= n; i++ {
		pos = (pos + direcao + n) % n
		if jogadores[pos] != 0 {
			return pos
		}
	}
	return pos
}

func main() {
	N, E, F := 0, 0, 0
	fmt.Scan(&N, &E, &F)

	jog := make([]int, 0, N)
	sinal := F

	for i := 1; i <= N; i++ {
		jog = append(jog, i*sinal)
		sinal *= -1
	}
	E -= 1
	vivos := N

	for vivos > 1 {
		print_jog(jog, E)
		direcao := 1
		if jog[E] < 0 {
			direcao = -1
		}

		target := procura_vivo(jog, E, direcao)
		jog[target] = 0
		vivos--

		E = procura_vivo(jog, E, direcao)
	}
	print_jog(jog, E)
}
