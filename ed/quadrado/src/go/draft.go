package main
import "fmt"

func calcQuadrado(n int) int {
    return (n-1) * (n-1) + (n-1) * 2 + 1 
}

func quadradoVoid(n int, inicial int, m bool, sqr int) {
    if n > inicial {
        return
    }

    if n == 1 {
        fmt.Printf("%d^2 = 1\n", n)
        sqr = n
        quadradoVoid(n + 1, inicial, m, sqr)
    }

    if n > 1 && sqr != 0 {
        sqr = calcQuadrado(n)
        fmt.Printf("%d^2 = %d^2 + 2*%d + 1 = %d\n", n, n-1, n-1, sqr)
        quadradoVoid(n + 1, inicial, m, sqr)
        return
    }
    
    if n > 1 && sqr == 0 {
        fmt.Printf("%d^2 = %d^2 + 2*%d + 1 = ?\n", n, n-1, n-1)
        quadradoVoid(n-1, inicial, m, sqr)
    }
}

func quadrado(n int) int {
    if n == 1 {
        fmt.Println("1^2 = 1")
        return 1
    }

    fmt.Printf("%d^2 = %d^2 + 2*%d + 1 = ?\n", n, n-1, n-1)
    
    result := quadrado(n - 1) + 2 * (n-1) + 1
    
    fmt.Printf("%d^2 = %d^2 + 2*%d + 1 = %d\n", n, n-1, n-1, result)

    return result
}

func main() {
    var n int
    fmt.Scan(&n)

    // var m bool
    // if n > 1 {
    //     m = true
    // } else {
    //     m = false
    // }

    // var sqr int = 0
    // quadradoVoid(n, n, m, sqr)

    quadrado(n)
}
