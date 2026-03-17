package main
import "fmt"

type Jogador struct {
    indice int
    pedras [2]int
    classificado bool
    pontuacao int
}

func main() {
    var n int
    fmt.Scan(&n)
    vetorJog := make([]Jogador, n)

    for i := 0; i < n; i++ {
        vetorJog[i].indice = i
        fmt.Scan(&vetorJog[i].pedras[0])
        fmt.Scan(&vetorJog[i].pedras[1])
        if (vetorJog[i].pedras[0] < 10 || vetorJog[i].pedras[1] < 10) {
            vetorJog[i].classificado = false;
        } else  {
            vetorJog[i].classificado = true;    
        }
        var temp int
        temp = vetorJog[i].pedras[0] - vetorJog[i].pedras[1]
        if temp < 0 {
            vetorJog[i].pontuacao = (temp * -1)
        } else {
            vetorJog[i].pontuacao = temp
        }
    }

    var maiorPont, maiorInd, nDes int
    nDes = 0
    for i := 0; i < n; i++ {
        if !vetorJog[i].classificado {
            nDes++
        }
        if vetorJog[i].classificado {
            maiorPont = vetorJog[i].pontuacao
            maiorInd = vetorJog[i].indice
        }
    }

    for i := 0; i < n; i++ {
        if !vetorJog[i].classificado {
            continue
        }
        if maiorPont > vetorJog[i].pontuacao {
            maiorPont = vetorJog[i].pontuacao
            maiorInd = vetorJog[i].indice
        }
    }

    if nDes == n {
        fmt.Printf("sem ganhador\n")
    } else {   
        fmt.Printf("%d\n", maiorInd)
    }
}
