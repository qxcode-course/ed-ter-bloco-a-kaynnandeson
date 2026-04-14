package main

import "fmt"

func main() {
	var total, baruel int
	fmt.Scan(&total, &baruel)

	vet := make([]int, total)
	for i := 0; i < total; i++ {
		vet[i] = i + 1
	}

	fig := make([]int, baruel)
	rep := make([]int, baruel)
	for i := 0; i < baruel; i++ {
		fmt.Scan(&fig[i])
		if i > 0 && fig[i] == fig[i-1] {
			rep[i] = fig[i]
		}
		for j := 0; j < total; j++ {
			if fig[i] == vet[j] {
				vet[j] = 0
				break
			}
		}
	}

	first := true
	vazio := true
	for i := 0; i < baruel; i++ {
		if rep[i] != 0 {
			if first {
				fmt.Print(rep[i])
				first = false
			} else {
				fmt.Print(" ", rep[i])
			}
			vazio = false
		}
	}
	if vazio {
		fmt.Print("N")
	}

	fmt.Println()

	first = true
	vazio = true
	for i := 0; i < total; i++ {
		if vet[i] != 0 {
			if first {
				fmt.Print(vet[i])
                first = false
			} else {
				fmt.Print(" ", vet[i])
			}
			vazio = false
		}
	}
	if vazio {
		fmt.Print("N")
	}
    fmt.Println()
}
