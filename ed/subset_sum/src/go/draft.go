package main

import "fmt"

func subset_sum(conjunto []int, k int, tam int, soma int) bool {
	if soma == k {
		return true
	}

	if tam >= len(conjunto) || soma > k {
		return false
	}

	if subset_sum(conjunto, k, tam+1, soma+conjunto[tam]) {
		return true
	}

    if subset_sum(conjunto, k, tam+1, soma) {
		return true
	}

	return false
}

func main() {
	n, k := 0, 0
	fmt.Scan(&n, &k)

	conjunto := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&conjunto[i])
	}

	if subset_sum(conjunto, k, 0, 0) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}
