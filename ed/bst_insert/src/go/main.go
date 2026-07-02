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
	Left  *Node
	Right *Node
}


func BstInsert(values []int) *Node {
	var no *Node
	for _, v := range values {
		no = insert(no, v)
	}
	return no
}

func insert(node *Node, value int) *Node {
	if node == nil { // criar novo no
		return &Node{
			Value: value,
		}
	}

	if value == node.Value { // se ja ta nao entra
		return node
	}
	// colocar no lugar
	if value < node.Value {
		node.Left = insert(node.Left, value)
	} else {
		node.Right = insert(node.Right, value)
	}
	return node
}

// -----------------------------------------------------------------------------------
func BShow(node *Node, history string) {
	if node != nil && (node.Left != nil || node.Right != nil) {
		BShow(node.Left, history+"l")
	}
	for i := 0; i < len(history)-1; i++ {
		if history[i] != history[i+1] {
			fmt.Print("│   ")
		} else {
			fmt.Print("    ")
		}
	}
	if history != "" {
		if history[len(history)-1] == 'l' {
			fmt.Print("╭───")
		} else {
			fmt.Print("╰───")
		}
	}
	if node == nil {
		fmt.Println("#")
		return
	}
	fmt.Println(node.Value)
	if node.Left != nil || node.Right != nil {
		BShow(node.Right, history+"r")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	values := make([]int, 0, len(parts))
	for _, elem := range parts {
		v, err := strconv.Atoi(elem)
		if err == nil {
			values = append(values, v)
		}
	}
	root := BstInsert(values)
	BShow(root, "") // Chama a função de impressão formatada
}
