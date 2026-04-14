package main

import "fmt"

func main() {
	var qtd_album, qtd_possui int
	fmt.Scan(&qtd_album, &qtd_possui)

	figurinhas := make([]int, qtd_possui) 
	for i := range figurinhas { //range pra ler 
		fmt.Scan(&figurinhas[i])
	}

	unicos := make(map[int]bool) // vericação map[int]bool
	repetidos := make([]int, 0, qtd_possui) //slice com 0

	for _, figura := range figurinhas {//figura assume o valor em figura
		if unicos[figura] { // se sim, adiciona no no slice
			repetidos = append(repetidos, figura)
		} else {
			unicos[figura] = true
		}
	}

	saida := fmt.Sprintf("%v", repetidos) // imprime
	if saida == "[]" { // caso vazio, N
		fmt.Println("N")
	} else {
		fmt.Println(saida[1 : len(saida)-1]) // imprime sem []
	}

	faltantes := make([]int, 0, qtd_album) // slice vazio pro que falta
	for i := 1; i <= qtd_album; i++ { // mesma verificação, só que ja verificada
		if !unicos[i] {
			faltantes = append(faltantes, i) // adiciona o numero das repetidas
		}
	}

	saida_f := fmt.Sprintf("%v", faltantes) // mesmo esquema de impressao
	if saida_f == "[]" {
		fmt.Println("N")
	} else {
		fmt.Println(saida_f[1 : len(saida_f)-1])
	}
}
