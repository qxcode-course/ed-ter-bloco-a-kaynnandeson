package main
import "fmt"

func diagonal(p string, e int) {
    tam := len(p)
    if (e == tam) {
        return
    }
    
    for i := 0; i < e; i++ {
        fmt.Print(" ")
    }
    
    fmt.Printf("%c\n", p[e])
    
    diagonal(p, e+1)
}

func main() {
    var palavra string
    fmt.Scan(&palavra)

    diagonal(palavra, 0)
}
