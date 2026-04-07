package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)

	vetorN := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&vetorN[i])
	}
	
	var m int
	fmt.Scan(&m)

	vetorM := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&vetorM[i])
		for j := 0; j < n; j++ {
			if (vetorM[i] == vetorN[j]) {
				vetorN[j] = 0
			}
		}
	}

	for i := 0; i < n; i++ {
		if vetorN[i] != 0 {
			fmt.Printf("%d ", vetorN[i])
		}
	}
	fmt.Println()
}
