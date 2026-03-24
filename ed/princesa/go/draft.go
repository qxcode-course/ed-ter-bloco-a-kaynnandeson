package main
import "fmt"
func main() {
    var n, e int
    fmt.Scan(&n, &e)

    vetor := make([]int, n)
    for i := 0; i < n; i++ {
        vetor[i] = i + 1;
    }
    
    pos := 0
    for i := 0; i < n; i++ {
        if vetor[i] == e {
            pos = i
            break
        }
    }

    countElem := n
    for countElem > 1 {
        fmt.Printf("[ ")
        for i := 0; i < n; i++ {
            if vetor[i] != 0  {
                fmt.Printf("%d", vetor[i])
            }
            if i == pos {
                fmt.Printf(">")
            }
            if vetor[i] != 0 {
                fmt.Printf(" ")
            }
        }
        fmt.Printf("]\n")
        
        prox := pos
        for {
            prox = (prox + 1) % n
            if vetor[prox] != 0 {
                break
            }
        }
        vetor[prox] = 0

        for {
            pos = (pos + 1) % n
            if vetor[pos] != 0 {
                break
            }
        }
        countElem--
    }
    
    fmt.Printf("[ ")
    for i := 0; i < n; i++ {
        if vetor[i] != 0  {
            fmt.Printf("%d", vetor[i])
        }
        if i == pos {
            fmt.Printf(">")
        }
        if vetor[i] != 0 {
            fmt.Printf(" ")
        }
    }
    fmt.Printf("]\n")
}
