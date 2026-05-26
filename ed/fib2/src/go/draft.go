package main
import "fmt"

var memo [41]int

func f(n int) int {
	if memo[n] != 0 {
		return memo[n]
	}

	if n == 1 || n == 2 {
		return 1
	}

	if n == 3 || n == 4 {
		return 2
	}

	memo[n] = f(n-2) + f(n-3)

	return memo[n]
}

func main() {
    var n int
    fmt.Scan(&n)

    fmt.Println(f(n))
}
