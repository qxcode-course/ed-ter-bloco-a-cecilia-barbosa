package main

import "fmt"

func C(n, k int) int {
	if k == 0 || k == n {
		return 1
	}
	return C(n-1, k-1) + C(n-1, k)
}

func main() {
	n, k := 0, 0

	fmt.Scan(&n, &k)

	fmt.Println(C(n, k))
}
