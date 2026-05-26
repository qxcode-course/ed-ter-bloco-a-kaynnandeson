package main
import "fmt"

func fn(n, k int) int {
    if n == 1 || n == 2 {
        return 1
    }

    return fn(n-1, k) + k*fn(n-2, k)
}

func main() {
    var n, k int
    fmt.Scan(&n, &k)

    fmt.Println(fn(n, k))
}
