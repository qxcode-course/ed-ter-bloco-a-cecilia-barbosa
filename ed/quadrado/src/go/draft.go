package main
import "fmt"

func Quadrado(n int) int {
    if n == 1 {
        fmt.Println("1^2 = 1")
        return 1
    }

    fmt.Printf("%d^2 = %d^2 + 2*%d + 1 = ?\n", n, n-1, n-1)

    anterior := Quadrado(n-1)
    result := anterior + 2*(n-1) + 1

    fmt.Printf("%d^2 = %d^2 + 2*%d + 1 = %d\n", n, n-1, n-1, result)
    return result
}

func main() {
    n := 0

    fmt.Scan(&n)

    Quadrado(n)
}