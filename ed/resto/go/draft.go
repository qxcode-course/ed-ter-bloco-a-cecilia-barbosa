package main

import "fmt"

func rec(num int) {
	if num == 0 {
		return 
	} else {
        div := num / 2
	    resto := num % 2
	    rec(div)
        fmt.Printf("%v %v\n", div, resto)
    }
}

func main() {
	num := 0
	fmt.Scan(&num)
	rec(num)
}
