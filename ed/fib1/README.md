package main
import "fmt"

func fn(n int) int {
    return fn(n-1) + fn(n-2)
}

func main() {
    var n, m int
    fmt.Scan(&n, &m)

    fmt.Println(fn(n))
}
