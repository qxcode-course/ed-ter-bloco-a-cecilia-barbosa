package main
import "fmt"
func main() {
    var Q int
    fmt.Scan(&Q)

    var D rune
    fmt.Scan(&D)

    x := make([]int, Q) //direita e esquerda
    y := make([]int, Q)

    for i:=0; i < Q; i++{
        fmt.Scan(&x[i], &y[i])
    }
    
    if D =='R'{
        for i:=0; i < Q; i++{
            fmt.Printf("%v %v\n", x[i]+1, y[i])
        }
    } 
    
    if D =='L'{
        for i:=0; i < Q; i++{
            fmt.Printf("%v %v\n", x[i]-1, y[i])
        }
    } 
    
    if D =='U'{
        for i:=0; i < Q; i++{
            fmt.Printf("%v %v\n", x[i], y[i]-1)
        }
    } 
    
    if D =='D'{
        for i:=0; i < Q; i++{
            fmt.Printf("%v %v\n", x[i]+1, y[i]+1)
        }
    } 

}
 