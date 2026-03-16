package main

import "fmt"

func main() {
	var N, M int
	rep := -1
	faltam := -1

	fmt.Scan(&N) // total do album
	fmt.Scan(&M) // quantas baruel tem

	v := make([]int, M)

	for i := 0; i < M; i++ {
		fmt.Scan(&v[i])
	}

	var repetidos [100]int
	rtam := 0

	for i := 0; i < M; i++ {
		for j := 0; j < M; j++ {
			if v[i] == v[j] {
				rep = 1
				rtam++
			}
		}
	}

	if rep == -1 {
		fmt.Printf("N\n")
	} else {
		for i := 0; i < rtam; i++{

		}
	}

	if faltam == -1 {
		fmt.Printf("N\n")
	} else {
		for i := 0; i < rtam; i++{
			
		}
	}

}
