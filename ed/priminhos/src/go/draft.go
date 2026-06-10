package main
import "fmt"

func ehPrimo(n int) bool {
    if n < 2 {
        return false
    }

    for i := 2; i*i <= n; i++ {
        if n % i == 0 {
            return false
        }
    }

    return true
}

func primos(n int) []int {
    var vet []int

    for i := 2; n > 0; i++ {
        if ehPrimo(i) {
            vet = append(vet, i)
            n--
        }
    }

    return vet
}

func main() {
    var n int
    fmt.Scan(&n)

    vet := primos(n)

    fmt.Printf("[")
    for i := 0; i < len(vet); i++ {
        if i == len(vet) - 1 {
            fmt.Print(vet[i])
            continue
        }
        fmt.Print(vet[i], ", ")
    }
    fmt.Println("]")
}
