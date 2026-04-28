package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MyList struct { // é a struct que vai ser usada pro slice la do inicio da main nascer
	data []int // slice de inteiros
}

func NewMyList(values []int) *MyList { // funciona como construtor
	return &MyList{ // cria uma lista nova com os valores que já existiam
		data: values} // aí retorna a nova struct que sera modificada
} 


type Iterator struct { // guarda a posição atual durante a execução, só caminha
	data  []int // dados
	index int // ponto de começo e parada
}


func (l *MyList) Iterator() *Iterator { // pertence ao tipo mylist e devolve o endereço do iteratpr
	return &Iterator{
		data:  l.data, // dado do tipo my list
		index: -1} // -1 pra começar do 0
}

func (i *Iterator) HasNext() bool { // checa se há outro
	return i.index < len(i.data)-1   // retorna true essa condição for verdade
}

func (i *Iterator) Next() int {
	if i.index == len(i.data) { // verifica se chegou no final
		panic(fmt.Errorf("No more elements"))
	}
	i.index += 1 // se nao chegar no final, acrescenta
	return i.data[i.index] // retorna elemento atual
}

type ReverseIterator struct {
	data  []int
	index int
}

func (l *MyList) ReverseIterator() *ReverseIterator {
	return &ReverseIterator{
		data:  l.data,
		index: len(l.data)}
}

func (i *ReverseIterator) HasNext() bool {
	return i.index > 0 // true se a condição, false se nao for
}

func (i *ReverseIterator) Next() int {
    i.index -= 1

	return i.data[i.index]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	mylist := NewMyList([]int{})
	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			break
		case "read":
			for i := 1; i < len(args); i++ {
				slice := make([]int, len(args)-1)
				for i, value := range args[1:] {
					slice[i], _ = strconv.Atoi(value)
				}
				mylist = NewMyList(slice)
			}
		case "show":
			fmt.Print("[ ")
			for it := mylist.Iterator(); it.HasNext(); {
				fmt.Printf("%v ", it.Next())
			}
			fmt.Println("]")
		case "reverse":
			fmt.Print("[ ")
			for it := mylist.ReverseIterator(); it.HasNext(); {
			fmt.Printf("%v ", it.Next())
			}
			fmt.Println("]")
		case "cyclic":
			// qtd, _ := strconv.Atoi(args[1])
			// fmt.Print("[ ")
			// it := mylist.CyclicIterator()
			// for range qtd {
			// 	fmt.Printf("%v ", it.Next())
			// }
			// fmt.Println("]")
		}
	}

}
