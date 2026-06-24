package main

import "fmt"

// x: número que está sendo testado
// div: divisor que está sendo testado
func priminhos(x int, div int, vet []int, n int) []int {
	if len(vet) == n {
		return vet
	}

	if div == x {
		vet = append(vet, x)
		return priminhos(x+1, 2, vet, n)
	}

	if x < 2 || x%div == 0{
		return priminhos(x+1, 2, vet, n)
	}

	return priminhos(x, div+1, vet, n)
}

func main() {
	n := 0
	fmt.Scan(&n)

	result := priminhos(2, 2, []int{}, n)

	fmt.Print("[")
	for i := 0; i < len(result); i++ {
		fmt.Print(result[i])
		if i < len(result)-1 {
			fmt.Print(", ")
		}
	}

	fmt.Println("]")
}
