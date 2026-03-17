package main
import "fmt"
func main() {
    var h, p, f, d int;

    fmt.Scan(&h, &p, &f, &d)

    if d == 1 {
        for i := f; i <= 15; i++ {
            if i == p {
                fmt.Println("N")
                break
            }
            if i == h {
                fmt.Println("S")
                break
            }
            if i == 15 {
                i = -1
            }
        }
    }
    if d == -1 {
        for i := f; i >= 0; i-- {
            if i == p {
                fmt.Println("N")
                break
            }
            if i == h {
                fmt.Println("S")
                break
            }
            if i == 0 {
                i = 16
            }
        }
    }
}
