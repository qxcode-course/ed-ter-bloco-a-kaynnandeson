package main
import "fmt"

func imprimir(p string, n int) {
    fmt.Printf("%s\n", p[len(p) - n:])
    
    tam := len(p)
    if tam == n {
        return
    }
    
    imprimir(p, n + 1)
}

func main() {
    var palavra string
    fmt.Scan(&palavra)

    imprimir(palavra, 1)
}
