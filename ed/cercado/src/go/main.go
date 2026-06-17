package main

import (
	"bufio"
	"fmt"
	"os"
)

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	lin := len(board)
	col := len(board[0]) 

	visitado := make([][]bool, lin)
	for i := range visitado {
		visitado[i] = make([]bool, col)
	}

	var dfs func(int, int)

	dfs = func(l, c int) {
		if l < 0 || l >= lin || c < 0 || c >= col {
			return
		} 

		if board[l][c] != 'O' || visitado[l][c] {
			return
		}

		visitado[l][c] = true

		dfs(l+1, c)
		dfs(l-1, c)
		dfs(l, c+1)
		dfs(l, c-1)
	}

	for l := 0; l < lin; l++ {
		dfs(l, 0)
		dfs(l, col-1)
	}

	for c := 0; c < col; c++ {
		dfs(0, c)
		dfs(lin-1, c)
	}

	for l := 0; l < lin; l++ {
		for c := 0; c < col; c++ {
			if board[l][c] == 'O' && !visitado[l][c] {
				board[l][c] = 'X'
			}
		}
	}
}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
