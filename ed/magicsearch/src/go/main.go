package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MagicSearch(slice []int, value int) int {
	_, _ = slice, value
	inicio, fim := 0, len(slice)-1

	
	for inicio <= fim {
		meio := (inicio + fim) / 2 // divide em dois pra binariedade

		if slice[meio] == value {
			inicio = meio+1 // empurra ate o ultimo encontrado 
		} else if slice[meio] > value { // se o nuemro do meio for maior que valro,
			fim = meio - 1 //fim vai pro inicio
		} else if slice[meio] < value { // se o numero do meio for menor que valor,
			inicio = meio + 1 // inicio vai pro fim
		}
	}
	 
	if inicio > 0 && slice[inicio-1] == value{ // diminui 1 pq ele empurra ate quando encontra o ultimo
		return inicio-1
	}else {
		return inicio
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	slice := make([]int, 0, 1)
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}

	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	result := MagicSearch(slice, value)
	fmt.Println(result)
}
