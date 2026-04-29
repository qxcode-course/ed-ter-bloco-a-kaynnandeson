package main
import "fmt"

func dividir(n int) {
    if (n == 0) {
        return 
    }

    dividir(n / 2)
    fmt.Println(n/2, n%2)
}

func main() {
    var n int
    fmt.Scan(&n)

    dividir(n)
}
