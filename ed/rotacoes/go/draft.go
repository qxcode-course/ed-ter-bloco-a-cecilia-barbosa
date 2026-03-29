package main

import "fmt"

func main() {
	tamT := 0
	rot := 0
	fmt.Scan(&tamT, &rot)

	vetor := make([]int, tamT)

	for i := 0; i < tamT; i++ {
		fmt.Scan(&vetor[i])
	}

	if rot == 0 {
		fmt.Printf("[")
		for _, i := range vetor {
			fmt.Printf(" %v", i)
		}
		fmt.Println(" ]")
	} else {
		for i := 1; i <= rot; i++ {
			j := 0
			vetor[j] = vetor[len(vetor)-1]
			j++

		}

		fmt.Printf("[")
		for _, i := range vetor {
			fmt.Printf(" %v", i)
		}
		fmt.Println(" ]")
	}
}
