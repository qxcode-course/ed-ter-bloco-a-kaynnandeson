package main
import "fmt"

func eh_primo(x int, div int) bool {
	if x < 2 {
		return false
	}
	if div*div > x {
		return true
	}
	if x%div == 0 {
		return false
	}
	return eh_primo(x, div+1)
}

func main() {
    var n int
    fmt.Scan(&n)

    primo := 1
    for n > 0 {
        if eh_primo(primo + 1, 2) {
            n--
        } 
        primo++
    }

    fmt.Println(primo)
}
