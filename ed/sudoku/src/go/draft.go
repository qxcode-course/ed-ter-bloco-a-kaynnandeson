package main
import "fmt"

func existeLinha(matriz [][]rune, lin int, valor rune) bool {
    for c := 0; c < len(matriz); c++ {
        if matriz[lin][c] == valor {
            return true
        }
    }
    return false
}

func existeColuna(matriz [][]rune, col int, valor rune) bool {
    for l := 0; l < len(matriz); l++ {
        if matriz[l][col] == valor {
            return true
        }
    }
    return false
}

func tamanhoQuadrante(n int) int {
    if n ==  4 {
        return 2
    }
    return 3
}

func existeQuadrante(matriz [][]rune, lin, col int, valor rune) bool {
    tam := tamanhoQuadrante(len(matriz))

    inicioLin := (lin / tam) * tam
    inicioCol := (col / tam) * tam

    for l := inicioLin; l < inicioLin + tam; l++ {
        for c := inicioCol; c < inicioCol + tam; c++ {
            if matriz[l][c] == valor {
                return true
            }
        }
    }

    return false
}

func valido(matriz [][]rune, lin, col int, valor rune) bool {
    if existeLinha(matriz, lin, valor) {
        return false
    }

    if existeColuna(matriz, col, valor) {
        return false
    }

    if existeQuadrante(matriz, lin, col, valor) {
        return false
    }

    return true
}

func resolver(matriz [][]rune, index int) bool {
    n := len(matriz)

    if index == n * n {
        return true
    }

    l := index / n
    c := index % n

    if matriz[l][c] != '.' {
        return resolver(matriz, index + 1)
    }

    for num := 1; num <= n; num++ {
        valor := rune('0' + num) 

        if valido(matriz, l, c, valor) {
            matriz[l][c] = valor

            if resolver(matriz, index+1) {
                return true
            }

            matriz[l][c] = '.'
        }
    }

    return false
}

func sudoku(matriz [][]rune) [][]rune {
    // for i := 1; i <= n; i++ {
    //     for l := 0; l < n; l++ {
    //         for c := 0; c < n; c++ {
    //             if matriz[l][c] != '.' {
    //                 continue
    //             }
    //             if valido(matriz, l, c, rune('0' + i)) {
    //                 matriz[l][c] = rune(i)
    //             }
    //         }
    //     }
    // }

    resolver(matriz, 0)

    return matriz
}

func main() {
    var n int
    fmt.Scan(&n)

    matriz := make([][]rune, n)
    for i := 0; i < n; i++ {
        var linha string
        fmt.Scan(&linha)

        matriz[i] = []rune(linha)
    }

    matrizResolvida := sudoku(matriz)
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            fmt.Printf("%c", matrizResolvida[i][j])
        }
        fmt.Println()
    }
}