package main
import "fmt"

type Pos struct {
    x, y int
}

func main() {
    var q int
    var d string
    fmt.Scan(&q, &d)

    vetor := make([]Pos, q)
    for i := 0; i < q; i++ {
        fmt.Scan(&vetor[i].x, &vetor[i].y)
    }

    if d == "L" {
        for i := q - 1; i > 0; i-- {
            vetor[i].x = vetor[i - 1].x
            vetor[i].y = vetor[i - 1].y
        }
        vetor[0].x = vetor[0].x - 1
    }
    if d == "R" {
        for i := q - 1; i > 0; i-- {
            vetor[i].x = vetor[i - 1].x
            vetor[i].y = vetor[i - 1].y
        }
        vetor[0].x = vetor[0].x + 1
    }
    if d == "U" {
        for i := q - 1; i > 0; i-- {
            vetor[i].x = vetor[i - 1].x
            vetor[i].y = vetor[i - 1].y
        }
        vetor[0].y = vetor[0].y - 1
    }
    if d == "D" {
        for i := q - 1; i > 0; i-- {
            vetor[i].x = vetor[i - 1].x
            vetor[i].y = vetor[i - 1].y
        }
        vetor[0].y = vetor[0].y + 1
    }

    for i := 0; i < q; i++ {
        fmt.Printf("%d %d", vetor[i].x, vetor[i].y)
        fmt.Println()
    }
}
