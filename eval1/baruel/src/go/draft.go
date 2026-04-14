package main

import "fmt"

func main() {
	qtd_album := 0
	fmt.Scan(&qtd_album)

	qtd_fig := 0
	fmt.Scan(&qtd_fig)

	baruel := make([]int, qtd_fig)

	for i := 0; i < len(baruel); i++ {
		fmt.Scan(&baruel[i])
	}

    unicas := make(map[int]bool)
    
    for rep := range baruel {
        if unicas[rep]{
            
        } else {
            
        }
    }

    for rep := range baruel {
        if unicas[rep]{
            
        } else {
            
        }
    }

}
