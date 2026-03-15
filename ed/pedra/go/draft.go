package main

import "fmt"

func main() {
	var N, i, A, B, diff int
	fmt.Scan(&N)

	v := make([]int, N)

	menor := 1000
	ganhador := -1

	for i = 0; i < N; i++ {
		fmt.Scan(&A, &B)

		if A >= 10 && B >= 10 {
			diff = A - B
			// evitar valor negativo
			if diff < 0 {
				diff *= -1
				v[i] = diff
			} else {
				v[i] = diff
			}

			if v[i] < menor {
				menor = v[i]
				ganhador = i
			}
		} else {
			v[i] = 100
		}
	}

	if ganhador == -1 {
		fmt.Printf("sem ganhador\n")
	} else {
		fmt.Printf("%v\n", ganhador)
	}
}
