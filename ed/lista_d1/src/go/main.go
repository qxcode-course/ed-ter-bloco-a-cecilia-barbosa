package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int   // Valor é público
	next  *Node // o próximo nó da lista
	prev  *Node // nó anterior
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	dlist := &LList{}
	dlist.root = &Node{} // cria o nó
	dlist.root.next = dlist.root
	dlist.root.prev = dlist.root

	return dlist
}

func (l *LList) Size() int {
	return l.size
}

func (l *LList) Clear() {
	l.root.next = l.root
	l.root.prev = l.root
	l.size = 0
}

// inserindo no inicio
func (l *LList) PushFront(value int) {
	no := &Node{
		Value: value,
	}

	primeiro := l.root.next
	no.next = primeiro
	no.prev = l.root
	l.root.next = no
	primeiro.prev = no

	l.size++
}

// adiciona um novo nó com esse valor no fim da lista
func (l *LList) PushBack(value int) {
	no := &Node{
		Value: value,
	}

	final := l.root.prev
	no.prev = final
	no.next = l.root
	final.next = no
	l.root.prev = no

	l.size++
}

// remove o primeiro valor da lista se existir
func (l *LList) PopFront() {
	if l.size == 0 {
		return
	}
	primeiro := l.root.next
	prox := primeiro.next
	l.root.next = prox
	prox.prev = l.root
	l.size--
}

// remove o último valor da lista se existir
func (l *LList) PopBack() {
	if l.size == 0 {
		return
	}
	final := l.root.prev
	pFinal := final.prev
	pFinal.next = l.root
	l.root.prev = pFinal
	l.size--
}

// quase a msm de vetbuild
func (l *LList) String() string {
	if l.size == 0 {
		return "[]"
	}
	var result strings.Builder
	no := l.root.next
	fmt.Fprintf(&result, "%d", no.Value)
	no = no.next

	for i := 1; i < l.size; i++ {
		fmt.Fprintf(&result, ", %d", no.Value)
		no = no.next
	}
	return "[" + result.String() + "]"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
