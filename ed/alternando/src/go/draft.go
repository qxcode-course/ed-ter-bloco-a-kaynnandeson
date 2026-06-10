package main
import "fmt"

func abs(n int) int {
    if n >= 0 {
        return n
    }

    return -n
}

func printar(vet []int, e int) {
    fmt.Printf("[ ")
    for i := 0; i < len(vet); i++ {
        if e == i {
            if vet[i] > 0 {
                fmt.Print(vet[i], ">")
            }
            if vet[i] < 0 {
                fmt.Print("<", vet[i])
            }
        } else {
            fmt.Print(vet[i])
        }
        fmt.Printf(" ")
    }
    fmt.Printf("]\n")
}

func matar(vet []int, e int) ([]int, int) {
    var alvo int

    if vet[e] > 0 {
        alvo = (e + 1) % len(vet)
    } else {
        alvo = (e - 1 + len(vet)) % len(vet)
    }

    vet = append(vet[:alvo], vet[alvo+1:]...)

    if len(vet) == 0 {
        return vet, 0
    }

    if alvo < e {
        e--
    }

    e = (e + 1) % len(vet)

    return vet, e
}

func princesa(tam, e, f int) {
    var n int

    if f == 1 {
        n = 1
    }
    if f == -1 {
        n = -1
    }

    var vet []int

    for i := 0; i < tam; i++ {
        vet = append(vet, n)
        
        nN := abs(n) + 1
        if n > 0 {
            n = -nN
            continue
        }
        if n < 0 {
            n = nN
        }
    }

    e-- 
    printar(vet, e)

    for len(vet) > 1 {
        vet, e = matar(vet, e)
        printar(vet, e)
    }
}

func main() {
    var n, e, f int

    fmt.Scan(&n, &e, &f)

    princesa(n, e, f)
}
