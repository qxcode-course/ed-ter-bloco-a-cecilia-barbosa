package main

import "fmt"

func main() {
	var H, P, F, D int
	fmt.Scan(&H, &P, &F, &D)

	if F < P && F < H && P < H { // POLICIA ENTRE
		if D == -1 {
			fmt.Printf("S\n")
		} else {
			fmt.Printf("N\n")
		}
	} else if F > P && F > H && P > H { // POLICIA ENTRE
		if D == 1 {
			fmt.Printf("S\n")
		} else {
			fmt.Printf("N\n")
		}
	} else if F > P && F < H && P < H { // F ENTRE
		if D == 1 {
			fmt.Printf("S\n")
		} else {
			fmt.Printf("N\n")
		}
	} else if F < P && F > H && P > H { // F ENTRE
		if D == -1 {
			fmt.Printf("S\n")
		} else {
			fmt.Printf("N\n")
		}
	} else if F > P && F > H && P < H { // H ENTRE
		if D == -1 {
			fmt.Printf("S\n")
		} else {
			fmt.Printf("N\n")
		}
	} else if F > P && F < H && P > H { // H ENTRE
		if D == 1 {
			fmt.Printf("S\n")
		} else {
			fmt.Printf("N\n")
		}
	}

}
