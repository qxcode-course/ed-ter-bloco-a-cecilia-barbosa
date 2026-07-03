package main

import (
	"fmt"
)

func main() {
	queue := NewQueue[string]()
	filaTimes := "ABCDEFGHIJKLMNOP"

	for _, time := range filaTimes {
		queue.Enqueue(string(time))
	}

	for i := 0; i < 15; i++ {
		m, n := 0, 0
		fmt.Scan(&m, &n)

		t1 := queue.Dequeue()
		t2 := queue.Dequeue()
		if m > n {
			queue.Enqueue(t1)
		} else {
			queue.Enqueue(t2)
		}
	}
	fmt.Println(queue.Dequeue())
}
