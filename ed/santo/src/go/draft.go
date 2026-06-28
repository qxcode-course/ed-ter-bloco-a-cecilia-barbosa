package main
import "fmt"

func dinheiro(n int, c int, atual float64) float64 {
    if n == 0 {
        return atual
    }

    anterior := (atual + float64(c)) / 2
    return dinheiro(n-1, c, anterior)
}

func main() {
    n, c := 0, 0
    fmt.Scan(&n, &c)

    result := dinheiro(n, c, 0)

    fmt.Printf("%.2f\n", result)
}