package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BetterSearch(slice []int, value int) (bool, int) {
	_, _ = slice, value
	inicio, fim := 0, len(slice)-1

	for inicio <= fim { // vai repetir ate nao pode mais se fechar
		meio := (inicio+fim) / 2 // divide em dois pra binariedade

		if slice[meio] > value { // se o nuemro do meio for maior que valro, 
			fim = meio - 1 //fim vai pro inicio 
		} else if slice[meio] < value { // se o numero do meio for menor que valor,
			inicio = meio + 1 // inicio vai pro fim
		} else {
			return true, meio // fim do loop retorna o numero achado
		}
	}
	return false, inicio
}  
// nao existe, retorna onde deveria estar, vai retornar false, meio
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	slice := []int{}
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}
	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	found, result := BetterSearch(slice, value)
	if found {
		fmt.Println("V", result)
	} else {
		fmt.Println("F", result)
	}
}
