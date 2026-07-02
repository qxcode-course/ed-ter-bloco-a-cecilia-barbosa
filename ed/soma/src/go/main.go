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

// somar todos os nós
func rec_sum(node *Node) int {
	if node == nil {
		return 0
	}

	result := node.Value + rec_sum(node.Right) + rec_sum(node.Left)
	return result
}

// menor presente na arvore
func rec_min(node *Node) int {
	if node == nil {
		return 0
	}

	pontoP := node.Value

	if node.Left != nil {
		minEsq := rec_min(node.Left)
		if minEsq < pontoP {
			pontoP = minEsq
		}
	}

	if node.Right != nil {
		minDir := rec_min(node.Right)
		if minDir < pontoP {
			pontoP = minDir
		}
	}
	return pontoP
}

// MyShow imprime a árvore binária de forma formatada.
func MyShow(node *Node, nivel int) {
	_, _ = node, nivel
	if node == nil {
		return
	}

	// no nao tem filho
	if node.Left == nil && node.Right == nil {
		for i := 0; i < nivel; i++ {
			fmt.Print("....")
		}
		fmt.Println(node.Value)
		return
	}

	//esquerda
	if node.Left == nil {
		for i := 0; i < nivel+1; i++ {
			fmt.Print("....")
		}
		fmt.Println("#")
	} else {
		MyShow(node.Left, nivel+1)
	}

	// raiz
	for i := 0; i < nivel; i++ {
		fmt.Print("....")
	}
	fmt.Println(node.Value)

	//direita
	if node.Right == nil {
		for i := 0; i < nivel+1; i++ {
			fmt.Print("....")
		}
		fmt.Println("#")
	} else {
		MyShow(node.Right, nivel+1)
	}
}

func BShow(node *Node, heranca string) {
	if node != nil && (node.Left != nil || node.Right != nil) {
		BShow(node.Left, heranca+"l")
	}
	for i := 0; i < len(heranca)-1; i++ {
		if heranca[i] != heranca[i+1] {
			fmt.Print("│   ")
		} else {
			fmt.Print("    ")
		}
	}
	if heranca != "" {
		if heranca[len(heranca)-1] == 'l' {
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
		BShow(node.Right, heranca+"r")
	}
}

func create(parts *[]string) *Node {
	elem := (*parts)[0]
	*parts = (*parts)[1:]
	if elem == "#" {
		return nil
	}
	value, _ := strconv.Atoi(elem)
	node := &Node{Value: value}
	node.Left = create(parts)
	node.Right = create(parts)
	return node
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	root := create(&parts)
	fmt.Println("Arvore:")
	BShow(root, "")
	fmt.Printf("Soma: %d, Minimo: %d\n", rec_sum(root), rec_min(root))
}
