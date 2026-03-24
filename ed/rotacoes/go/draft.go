package main

import "fmt"

func main() {
	tamT := 0
	rot := 0
	fmt.Scan(&tamT, &rot)

	v := make([]int, tamT)

	for i := 0; i < tamT; i++ {
		fmt.Scan(&v[i])
	}

}
