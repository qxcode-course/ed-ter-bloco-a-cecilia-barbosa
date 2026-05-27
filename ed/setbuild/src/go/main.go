package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Set struct {
	data     []int
	size     int
	capacity int
}

func NewSet(capacity int) *Set {
	return &Set{
		data:     make([]int, 1),
		size:     0,
		capacity: 1,
	}
}

// msm de vetbuild
func (s *Set) Reserve(capNova int) { // identica a de vetbuild
	if capNova <= s.capacity {
		return
	}
	nData := make([]int, capNova)
	for i := range s.size {
		nData[i] = s.data[i]
	}
	s.data = nData
	s.capacity = capNova

}

// mesmo codigo de bettersearch
func (s *Set) binarySearch(value int) int {
	inicio, fim := 0, s.size-1

	for inicio <= fim { 
		meio := (inicio+fim) / 2 

		if s.data[meio] > value { 
			fim = meio - 1 
		} else if s.data[meio] < value { 
			inicio = meio + 1 
		} else {
			return meio 
		}
	}
	return -inicio-1
}  

// quase o mesmo de vetbuild tb
func (s *Set) Insert(value int) {
	indice := s.binarySearch(value)
	if indice >= 0 {
		return
	}

	posicao := -indice-1
	if s.size == s.capacity{
		s.Reserve(s.capacity*2)
	}

	for i := s.size; i > posicao; i--{
		s.data[i] = s.data[i-1]
	}

	s.data[posicao] = value
	s.size++
}

func (s *Set) Contains(value int) bool {
	return s.binarySearch(value) >= 0
}

// praticamente o mesmo de vetbuild
func (s *Set) Erase(value int) bool {
	indice := s.binarySearch(value)
	if indice < 0 {
		return false
	}

	for i := indice; i < s.size-1; i++{
		s.data[i] = s.data[i+1]
	}

	s.size--
	return true
}

//copiado de vetbuild
func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}

func (s *Set) Show() string {
	return "[" + Join(s.data[:s.size], ", ") + "]"
}



func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
			 	value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			fmt.Println(v.Show())
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			if !v.Erase(value){
				fmt.Println("value not found")
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value){
				fmt.Println("true")
			} else{
				fmt.Println("false")
			}
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
