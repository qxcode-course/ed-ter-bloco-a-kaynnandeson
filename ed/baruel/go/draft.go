package main
import "fmt"
func main() {
    var total, baruel int
    fmt.Scan(&total, &baruel)

    vetorTotal := make([]int, total)
    for i := 0; i < total; i++ {
        vetorTotal[i] = i + 1
    }

    vetorBaruel := make([]int, baruel)
    for i:= 0; i < baruel; i++ {
        fmt.Scan(&vetorBaruel[i])
        for j := 0; j < total; j++ {
            if (vetorBaruel[i] == vetorTotal[j]) {
                vetorTotal[j] = 0
            }
        }
    }

    var countRep int = 0
    var first bool = true
    for i := 0; i < baruel - 1; i++ {
        if vetorBaruel[i] == vetorBaruel[i + 1] {
            if !first {
                fmt.Print(" ")
            }
            fmt.Print(vetorBaruel[i])
            first = false
            countRep++
        } 
    }
    if countRep == 0 {
        fmt.Printf("N")
    }
    fmt.Println()

    var countCard int = 0
    first = true
    for i := 0; i < total; i++ {
        if vetorTotal[i] != 0 {
            if !first {
                fmt.Print(" ")
            }
            fmt.Print(vetorTotal[i])
            first = false
            countCard++
        }
    }
    if countCard == 0 {
        fmt.Printf("N")
    }
    fmt.Println()
}
