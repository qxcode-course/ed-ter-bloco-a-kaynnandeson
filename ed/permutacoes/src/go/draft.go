package main
import "fmt"
import "sort"

func gerar(s, atual string, usado []bool) {
    if len(atual) == len(s) {
        fmt.Println(atual)
        return
    }

    for i := 0; i < len(s); i++ {
        if usado[i] {
            continue
        }

        usado[i] = true

        gerar(s, atual + string(s[i]), usado)

        usado[i] = false
    }
}

func permutar(s string) {
    chars := []rune(s)

    sort.Slice(chars, func(i, j int) bool {
        return chars[i] < chars[j]
    })

    s = string(chars)

    usado := make([]bool, len(s))
    gerar(s, "", usado)
}

func main() {
    var s string

    fmt.Scan(&s)

    permutar(s)
}
