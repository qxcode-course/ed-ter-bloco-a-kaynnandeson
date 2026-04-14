package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	vetN := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&vetN[i])
	}
	
	var m int
	fmt.Scan(&m)
	vetM := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&vetM[i])
		for j := 0; j < n; j++ {
			if vetM[i] == vetN[j] {
				vetN[j] = 0
				break
			}
		}
	}

	for i := 0; i < n; i++ {
		if vetN[i] != 0 {
			fmt.Print(vetN[i], " ")
		}
	}
	fmt.Println()
}
