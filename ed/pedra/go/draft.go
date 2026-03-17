package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    vetor := make([]int, n * 2)

    for i := 0; i < n * 2; i++ {
        fmt.Scan(&vetor[i])
        fmt.Printf("%d ", vetor[i])
    }

}
