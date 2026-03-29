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
	//ajuda da Sophia Muniz
	rotacao := rot % tamT // usa o resto para se preocupar apenas com as que sobraram,
	if rotacao < 0 {      //nao as que dao loop completo e ja volta do zero
		rotacao += tamT
	}
	// foi essa lógica que ela me ajudou

	vetFinal := make([]int, 0, tamT) // slice dinamico

	for i := tamT - rotacao; i < tamT; i++ { // coloca os rotacionados no inicio
		vetFinal = append(vetFinal, vetor[i])
	}
	for i := 0; i < tamT-rotacao; i++ { // coloca os primeiros no final
		vetFinal = append(vetFinal, vetor[i])
	}

	fmt.Printf("[")
	for _, i := range vetFinal {
		fmt.Printf(" %v", i)
	}
	fmt.Println(" ]")

}
