package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Multiset struct { //
	data     []int
	size     int
	capacity int
}

func NewMultiSet(capacity int) *Multiset { // novo vetor
	return &Multiset{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

// mesmos de vetbuild
func (v *Multiset) Reserve(capNova int) {
	if capNova <= v.capacity {
		return
	}
	nData := make([]int, capNova)
	for i := range v.size {
		nData[i] = v.data[i]
	}
	v.data = nData
	v.capacity = capNova

}

// quase igual de vetbuild
func (v *Multiset) Insert(value int) {
	if v.size == v.capacity {
		capNova := v.capacity * 2
		if capNova == 0 {
			capNova = 1
		}
		v.Reserve(capNova)
	}

	v.data[v.size] = value
	v.size++

	for i := v.size - 1; i > 0 && v.data[i] < v.data[i-1]; i-- {
		temp := v.data[i]
		v.data[i] = v.data[i-1]
		v.data[i-1] = temp
	}
}

// mesma de vetbuild
func (v *Multiset) IndexOf(value int) int {
	for i := range v.size {
		if v.data[i] == value {
			return i
		}
	}
	return -1
}

// mesma de vetbuild
func (v *Multiset) Contains(value int) bool {
	return v.IndexOf(value) != -1
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

// mesmo de vetbuild
func (v *Multiset) String() string {
	return "[" + Join(v.data[:v.size], ", ") + "]"
}

// quase a mesma de vetbuild
func (v *Multiset) Erase(indice int) {
	if indice < 0 || indice >= v.size {
		return
	}

	for i := indice; i < v.size-1; i++ {
		v.data[i] = v.data[i+1]
	}

	v.size--
}

// mesma de vetbuild
func (v *Multiset) Clear() {
	v.size = 0
}

func (v *Multiset) Count(value int) int {
	count := 0
	for i := range v.size {
		if v.data[i] == value {
			count++
		}
	}
	return count
}

func (v *Multiset) Unique() int {
	if v.size == 0 {
		return 0
	}
	uni := 1
	for i := 1; i < v.size; i++ {
		if v.data[i] != v.data[i-1] {
			uni++
		}
	}

	return uni
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms.String())
		case "erase":
			value, _ := strconv.Atoi(args[1])
			indice := ms.IndexOf(value)
			if indice == -1 {
				fmt.Println("value not found")
			} else {
				ms.Erase(indice)
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			if ms.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
