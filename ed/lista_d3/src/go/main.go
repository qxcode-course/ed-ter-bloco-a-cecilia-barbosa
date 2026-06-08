package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	next  *Node
	prev  *Node
	root  *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	list := &LList{}
	list.root = &Node{root: nil}
	list.root.next = list.root
	list.root.prev = list.root
	list.root.root = list.root // nó sentinela aponta pra si mesmo
	return list
}

// peguei do lista_d2
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

func (l *LList) PushBack(value int) {
	l.insertBefore(l.root, value)
}

func (l *LList) insertBefore(mark *Node, value int) {
	n := &Node{
		Value: value,
		root:  l.root,
	}
	n.prev = mark.prev
	n.next = mark
	mark.prev.next = n
	mark.prev = n

	l.size++
}

// tem que percorrer ate o lugar certo pra inserir
func addsorted(l *LList, value int) {
	no := l.root.next

	for no.Value <= value && no != l.root {
		no = no.next
	}

	l.insertBefore(no, value)
}

// só inverter a lista agr
func reverse(l *LList) {
	if l.size <= 1 { // caso tenha 1 ou menos nao inverter
		return
	}

	no := l.root

	for i := 0; i <= l.size; i++ {
		no.next, no.prev = no.prev, no.next
		no = no.prev
		if no == l.root {
			break
		}
	}
}

func compare(l1, l2 *LList) bool {
	if l1.size != l2.size {
		return false
	}

	num1 := l1.root.next
	num2 := l2.root.next
	for num1 != l1.root && num2 != l2.root {
		if num1.Value != num2.Value {
			return false
		}
		num1 = num1.next
		num2 = num2.next
	}
	return true
}

func merge(l1, l2 *LList) *LList {
	result := NewLList()
	num1 := l1.root.next
	num2 := l2.root.next

	for num1 != l1.root || num2 != l2.root {
		if num2 == l2.root || (num1 != l1.root && num1.Value <= num2.Value) {
			result.PushBack(num1.Value)
			num1 = num1.next
		} else {
			result.PushBack(num2.Value)
			num2 = num2.next
		}
	}
	return result
}

func str2list(serial string) *LList {
	serial = serial[1 : len(serial)-1]
	ll := NewLList()
	if serial == "" {
		return ll
	}
	for _, p := range strings.Split(serial, ",") {
		value, _ := strconv.Atoi(strings.TrimSpace(p))
		ll.PushBack(value)
	}
	return ll
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

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
		case "compare":
			lla := str2list(args[1])
			llb := str2list(args[2])
			if compare(lla, llb) {
				fmt.Println("iguais")
			} else {
				fmt.Println("diferentes")
			}
		case "addsorted":
			lla := NewLList()
			for i := 1; i < len(args); i++ { // vai checar pra cada um ja
				value, _ := strconv.Atoi(args[i])
				addsorted(lla, value)
			}
			fmt.Println(lla.String())
		case "reverse":
			lla := str2list(args[1])
			reverse(lla)
			fmt.Println(lla.String())
		case "merge":
			lla := str2list(args[1])
			llb := str2list(args[2])
			merged := merge(lla, llb)
			fmt.Println(merged)
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
