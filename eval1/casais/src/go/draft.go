package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    vetor := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&vetor[i])
    }

    countCasais := 0
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if vetor[i] == -vetor[j] && vetor[i] != 0 {
                countCasais++
                vetor[i] = 0
                vetor[j] = 0
            }
        }
    }

    fmt.Println(countCasais)
}
