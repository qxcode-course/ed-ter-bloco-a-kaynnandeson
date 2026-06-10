package main
import "fmt"

func seq(n, m int) int {
    pontos := 1
    
    if m == 1  {
        return pontos
    }

    incremento := n - 1

    for i := 2; i <= m; i++ {
        pontos += incremento
        incremento += n - 2
    }
   
    return pontos
}

func main() {
    var n, m int
    fmt.Scan(&n, &m)

    fmt.Println(seq(n, m))
}
