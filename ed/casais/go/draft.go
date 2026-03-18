package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    
    vetor := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&vetor[i])
    }
    
    var nCasais int = 0 
    for i := 0;i < n; i++  {
        var comp int = vetor[i];
        for j := i; j < n; j++ {
            if comp == (-vetor[j]) && comp != 0 {
                nCasais++
                vetor[i] = 0
                vetor[j] = 0
                break
            }
        }
    }

    fmt.Println(nCasais)
}
