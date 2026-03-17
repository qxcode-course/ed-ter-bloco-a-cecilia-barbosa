package main

import "fmt"

func main() {
	var N int
	casal := 0
	fmt.Scan(&N)
	v := make([]int, N) // todos

	var f [100]int
	var m [100]int

	//contar machos e fêmeas
	ftam := 0
	mtam := 0
	for i := 0; i < N; i++ {
		fmt.Scan(&v[i])
		if v[i] < 0 {
			f[ftam] = v[i]
			ftam++
		} else {
			m[mtam] = v[i]
			mtam++
		}
	}

	// contar casais
	var casada [100]bool // impedir que case dnovo

	for j := 0; j < mtam; j++ {

		for k := 0; k < ftam; k++ {
			if casada[k] == false && m[j] == (-1*f[k]) {
				casal++
				casada[k] = true
				break
			} else {
				casada[k] = false
			}
		}
	}

	fmt.Printf("%v\n", casal)

}
