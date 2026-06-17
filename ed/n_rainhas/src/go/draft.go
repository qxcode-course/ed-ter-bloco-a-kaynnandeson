package main
import "fmt"

func resolver(linha, n int, col, dg1, dg2 []bool) int {
    if linha == n {
        return 1
    }

    total := 0

    for c := 0; c < n; c++ {
        d1 := linha - c + (n - 1)
        d2 := linha + c

        if col[c] || dg1[d1] || dg2[d2] {
            continue
        }

        col[c] = true
        dg1[d1] = true
        dg2[d2] = true

        total += resolver(linha + 1, n, col, dg1, dg2)

        col[c] = false
        dg1[d1] = false
        dg2[d2] = false
    }
    
    return total
}

func nRainhas(n int, tabuleiro [][]int) int {
    col := make([]bool, n)
    dg1 := make([]bool, 2 * n-1)
    dg2 := make([]bool, 2 * n-1)

    return resolver(0, n, col, dg1, dg2)
}

func main() {
    var n int
    fmt.Scan(&n)

    tabuleiro := make([][]int, n)

    fmt.Println(nRainhas(n, tabuleiro))
}