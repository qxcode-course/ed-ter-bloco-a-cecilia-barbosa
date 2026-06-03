package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// mesma da outra questao
type Node struct {
	Value int   // Valor é público
	next  *Node // o próximo nó da lista
	prev  *Node // nó anterior
	root  *Node
}

// retorna o próximo nó ou nulo, se o próximo é o root
func (n *Node) Next() *Node {
	if n.next == n.root {
		return nil
	}
	return n.next
}

// retorna o nó anterior ou nulo, se o anterior é o root
func (n *Node) Prev() *Node {
	if n.prev == n.root {
		return nil
	}
	return n.prev
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

// msm da outra questao
func (l *LList) Size() int {
	return l.size
}

// mesmo tb
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

func (l *LList) Clear() {
	l.root.next = l.root
	l.root.prev = l.root
	l.size = 0
}

func (l *LList) PushFront(value int) {
	no := &Node{
		Value: value,
		root:  l.root,
	}

	primeiro := l.root.next
	no.next = primeiro
	no.prev = l.root
	l.root.next = no
	primeiro.prev = no

	l.size++
}

func (l *LList) PushBack(value int) {
	no := &Node{
		Value: value,
		root:  l.root,
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

// retorna o primeiro nó válido da lista ou nulo
func (l *LList) Front() *Node {
	if l.size == 0 {
		return nil
	}
	return l.root.next
}

func (l *LList) Back() *Node {
	if l.size == 0 {
		return nil
	}
	return l.root.prev
}

// retorna o nó que contém a primeira ocorrência desse valor ou nulo
func (l *LList) Search(value int) *Node {
	for no := l.root.next; no != l.root; no = no.next {
		if no.Value == value {
			return no
		}
	}
	return nil
}

// insere um novo nó antes do nó passado por referência
func (l *LList) Insert(node *Node, value int) {
	if node == nil || node.root != l.root {
		return
	}

	novoNo := &Node{
		Value: value,
		root:  l.root,
	}

	anterior := node.prev

	novoNo.next = node
	novoNo.prev = anterior

	anterior.next = novoNo
	node.prev = novoNo

	l.size++
}

// remove o nó passado por referência retornando o nó que ficou no lugar dele ou nulo
// caso o nó não exista ou o próximo seja o último

func (l *LList) Remove(node *Node) *Node {
	if node == nil || node.root != l.root || node == l.root {
		return nil
	}
	next := node.next
	prev := node.prev

	prev.next = next
	next.prev = prev

	l.size--
	return next
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
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.Value = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
