package main
import "fmt"

func comb(n, k int) int {
    if k == 0 || k == n {
        return 1
    }

    return comb(n-1, k-1) + comb(n-1, k)
}

func main() {
    var n, k int;
    fmt.Scan(&n, &k);

    fmt.Println(comb(n, k))
}
