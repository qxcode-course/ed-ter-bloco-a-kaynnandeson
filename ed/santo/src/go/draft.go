package main
import "fmt"

func jm(n, c int) float64 {
    dinheiro := 0.0

    for i := 0; i < n; i++ {
        dinheiro = (dinheiro + float64(c)) / 2
    }

    return dinheiro
}

func main() {
    var n, c int

    fmt.Scan(&n, &c)

    fmt.Printf("%.2f\n", jm(n, c))
}
