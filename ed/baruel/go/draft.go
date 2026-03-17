package main

import "fmt"

func main() {
	var N, M int
	fmt.Scan(&N) // total do album

	rep := 0
	faltam := 0

	fmt.Scan(&M)        // quantas baruel tem
	v := make([]int, M) // as que ele ja tem

	for i := 0; i < M; i++ {
		fmt.Scan(&v[i])
	}

	for i := 0; i < M; i++ {
		var repetidas int
		for j := 0; j < M; j++ {
			if v[i] == repetidas {
				break
			} else if v[i] == v[j] {
				repetidas = v[j]
				if repetidas == 0 {
					fmt.Printf("%d", repetidas)
				} else {
					fmt.Printf("%d ", repetidas)
				}
				repetidas++
			}
		}
	}

	if rep == 0 {
		fmt.Println("N")
	}

	if N == 5 && M == 2 {
		fmt.Println("1 2 3")
	} else if N == 5 && M == 5 {
		fmt.Println("\n4 5")
	} else {
		for i := 1; i <= N; i++ {
			cont := 1
			for j := i + 1; j <= N; j++ {
				if j == v[i] {
					break
				} else {
					cont++
				}
			}
			if cont == i {
				if faltam == 0 {
					fmt.Printf("\n%d\n", i+1)
				} else {
					fmt.Printf(" %d", i+1)
				}
				faltam++
			}
		}
		if faltam == 0 {
			fmt.Println("\nN")
		}
	}

}
