package main

import "fmt"

// x: número que está sendo testado
// div: divisor que está sendo testado
// Usando código do eh_primo que ja fiz
func eh_primo(x int, div int) bool {
	if x < 2 {
		return false
	}

	if div == x {
		return true
	}

	if x%div == 0 {
		return false
	}

	div++
	return eh_primo(x, div)
}

func enesimo(num int, cont int, n int) int {
	if eh_primo(num, 2) {
		cont++
	}

	if cont == n {
		return num
	}

	return enesimo(num+1, cont, n)
}

func main() {
	x := 0
	fmt.Scan(&x)

	result := enesimo(2, 0, x)
	fmt.Printf("%d\n", result)
}
