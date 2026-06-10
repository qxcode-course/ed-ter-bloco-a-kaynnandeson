package main
import "fmt"
import "math"


func descobrirE(n, k int) int {
    if n <= k {
        return 1
    }

    return descobrirE(int(math.Ceil(float64(n)/2.0)), k) + descobrirE(int(math.Floor(float64(n)/2)), k)
}

func main() {
    var n, k int

    fmt.Scan(&n, &k)

    fmt.Println(descobrirE(n, k))
}
