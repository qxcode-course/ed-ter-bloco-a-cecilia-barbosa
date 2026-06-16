package main
import "fmt"

func main() {
    entrada := ""
    fmt.Scanln(&entrada)

    texto := []rune{}
    posicao := 0

    for c := range entrada {
        switch c {
        case '>':
            if posicao < len(texto){
                posicao++
            }
        case '<':
            if posicao > 0 {
                posicao--
            }
        case 'B':
            if posicao < len(texto){
                texto = append(texto[:posicao-1], texto[posicao:]...)
                posicao--
            }
        case 'D':
            if posicao < len(texto){
                texto = append(texto[:posicao], texto[posicao+1:]...)
            }
        case 'R':
        default:
        }
    }

    result := string(texto[:posicao]) + "|" + string(texto[posicao:])
    fmt.Println(result)

}
