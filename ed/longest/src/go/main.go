package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	lin := len(matrix)
	col := len(matrix[0])

	memo := make([][]int, lin)
	for i := range memo {
		memo[i] = make([]int, col)
	}

	dirs := [][]int {
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	var dfs func(int, int) int

	dfs = func(l, c int) int {
		if memo[l][c] != 0 {
			return memo[l][c]
		}

		best := 1

		for _, d := range dirs {
			nl := l + d[0]
			nc := c + d[1]

			if nl < 0 || nl >= lin || nc < 0 || nc >= col {
				continue
			}

			if matrix[nl][nc] > matrix[l][c] {
				length := 1 + dfs(nl, nc)

				if length > best {
					best = length
				}
			}
		}
		memo[l][c] = best
		return best
	}

	ans := 0

	for l := 0; l < lin; l++ {
		for c := 0; c < col; c++ {
			path := dfs(l, c)

			if path > ans {
				ans = path
			}
		}
	}

	return ans
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	parts := strings.Fields(scanner.Text())
	if len(parts) < 2 {
		return
	}
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	matrix := make([][]int, nl)
	for i := 0; i < nl; i++ {
		if !scanner.Scan() {
			return
		}
		tokens := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc && j < len(tokens); j++ {
			v, _ := strconv.Atoi(tokens[j])
			row[j] = v
		}
		matrix[i] = row
	}

	fmt.Println(longestIncreasingPath(matrix))
}
