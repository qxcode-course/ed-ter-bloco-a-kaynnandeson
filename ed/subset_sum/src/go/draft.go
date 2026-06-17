package main
import "fmt"

func backtracking(vet []int, i, soma, k int) bool {
    if soma == k {
        return true
    }

    if i == len(vet) || soma > k {
        return false
    }

    if backtracking(vet, i+1, soma+vet[i], k) {
        return true
    }

    if backtracking(vet, i+1, soma, k) {
        return true
    }

    return false
}

func verificarSoma(vet []int, k int) bool {
    return backtracking(vet, 0, 0, k)
}

func main() {
    var n, k int

    fmt.Scan(&n, &k)

    vet := make([]int, n)

    for i := 0; i < n; i++ {
        fmt.Scan(&vet[i])
    }

    if verificarSoma(vet, k) {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
}