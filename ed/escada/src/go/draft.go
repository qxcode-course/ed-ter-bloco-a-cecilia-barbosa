package main

import "fmt"

func escada(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	a := 1
	b := 1
	c := 2

	atual := 0

	for i := 4; i <= n; i++ {
		atual = c + a
		a = b
		b = c
		c = atual
	}
	return c
}

func main() {
	n := 0
	fmt.Scan(&n)

    fmt.Printf("%d\n", escada(n))
}
