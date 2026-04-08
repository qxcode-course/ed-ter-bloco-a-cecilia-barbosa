package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// imprime vetor
func imprime(vet []int) string {
	if len(vet) == 1 {
		return fmt.Sprintf("%v", vet[0])
	}

	saida := fmt.Sprintf("%v, ", vet[0])
	return saida + imprime(vet[1:])
}

// sai com parentese
func tostr(vet []int) string {
	_ = vet
	if len(vet) == 0 {
		return "[]"
	}

	return "[" + imprime(vet) + "]"
}

func imprimeRev(vet []int) string {
	if len(vet) == 1 {
		return fmt.Sprintf("%v", vet[0])
	}

	ultimo := len(vet) - 1
	atual := fmt.Sprintf("%v, ", vet[ultimo])

	return atual + imprimeRev(vet[:ultimo])
}

// formata ao inverso[5, 4, 3]
func tostrrev(vet []int) string {
	if len(vet) == 0 {
		return "[]"
	}
	return "[" + imprimeRev(vet) + "]"
}

func reverso(vet []int, i int, f int) {
	if i > f {
		return
	}
	vet[i], vet[f] = vet[f], vet[i]
	reverso(vet, i+1, i-1)
}

// reverse: inverte os elementos do slice
func reverse(vet []int) {
	_ = vet
	reverso(vet, 0, len(vet)-1)
}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	_ = vet
	if len(vet) == 0 {
		return 0
	}

	primeiro := vet[0]
	resto := vet[1:]
	return primeiro + sum(resto)
}

// mult: produto dos elementos do slice
func mult(vet []int) int {
	if len(vet) == 0 {
		return 1
	}

	primeiro := vet[0]
	resto := vet[1:]
	return primeiro * mult(resto)
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {
    if len(vet) == 0 {
        return -1
    }

    var rec func(i int, idxMenor int) int
    
    rec = func(i int, idxMenor int) int {
        if i == len(vet) {
            return idxMenor
        }

        if vet[i] < vet[idxMenor] {
            return rec(i+1, i)
        }
        
        return rec(i+1, idxMenor)
    }

    return rec(0, 0)
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
