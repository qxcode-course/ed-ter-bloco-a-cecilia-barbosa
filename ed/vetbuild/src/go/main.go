package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	data     []int
	size     int
	capacity int
}

func NewVector(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity), // nunca use len(data) ou cap(data) ou qq método do go de manipulação de array
		size:     0,
		capacity: capacity,
	}
}

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

func (v *Vector) Status() string {
	return fmt.Sprintf("size:%d capacity:%d", v.size, v.capacity)
}

func (v *Vector) Insert(indice int, value int) error {
	if v.size == v.capacity {
		capNova := v.capacity * 2
		if capNova == 0 {
			capNova = 1
		}
		v.Reserve(capNova)
	}

	for i := v.size; i > indice; i-- {
		v.data[i] = v.data[i-1]
	}

	v.data[indice] = value
	v.size++
	return nil
}

func (v *Vector) Show() string {
	return "[" + Join(v.data[:v.size], ", ") + "]"
}

func (v *Vector) Clear() {
	v.size = 0
}

func (v *Vector) Reserve(capNova int) {
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

func (v *Vector) PushBack(value int) {
	if v.size == v.capacity {
		capNova := v.capacity * 2
		if capNova == 0 {
			capNova = 1
		}
		v.Reserve(capNova)
	}
	v.data[v.size] = value
	v.size++
}

func (v *Vector) Contains(value int) bool {
	return v.IndexOf(value) != -1
}

func (v *Vector) IndexOf(value int) int {
	for i := range v.size {
		if v.data[i] == value {
			return i
		}

	}
	return -1
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewVector(0)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
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
			v = NewVector(value)
		case "push":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.PushBack(value)
			}
		case "show":
			fmt.Println(v.Show())
		case "status":
			fmt.Println(v.Status())
		case "pop":
			// err := v.PopBack()
			// if err != nil {
			// 	fmt.Println(err)
			// }
		case "insert":
			indice, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Insert(indice, value)
			if err != nil {
				fmt.Println(err)
			}
		case "erase":
			// index, _ := strconv.Atoi(parts[1])
			// err := v.Erase(index)
			// if err != nil {
			// 	fmt.Println(err)
			// }
		case "indexOf":
			value, _ := strconv.Atoi(parts[1])
			index := v.IndexOf(value)
			fmt.Println(index)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
			v.Clear()
		case "capacity":
			// fmt.Println(v.Capacity())
		case "get":
			// index, _ := strconv.Atoi(parts[1])
			// value, err := v.At(index)
			// if err != nil {
			// 	fmt.Println(err)
			// } else {
			// 	fmt.Println(value)
			// }
		case "set":
			// index, _ := strconv.Atoi(parts[1])
			// value, _ := strconv.Atoi(parts[2])
			// err := v.Set(index, value)
			// if err != nil {
			// 	fmt.Println(err)
			// }
			//
		case "reserve":
			newCapacity, _ := strconv.Atoi(parts[1])
			v.Reserve(newCapacity)
		case "slice":
			// start, _ := strconv.Atoi(parts[1])
			// end, _ := strconv.Atoi(parts[2])
			// slice := v.Slice(start, end)
			// fmt.Println(slice)
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
