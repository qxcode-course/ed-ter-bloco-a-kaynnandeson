package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tostr(vet []int) string {
	if len(vet) == 0 {
		return "[]"
	}

	var rec func(v []int) string
	rec = func(v []int) string {
		if len(v) == 1 {
			return fmt.Sprintf("%d", v[0])
		}
		return fmt.Sprintf("%d, %s", v[0], rec(v[1:]))
	}

	return "[" + rec(vet) + "]"
}

func tostrrev(vet []int) string {
	if len(vet) == 0 {
		return "[]"
	}

	var rec func(v []int) string
	rec = func(v []int) string {
		if len(v) == 1 {
			return fmt.Sprintf("%d", v[0])
		}
		return fmt.Sprintf("%d, %s", v[len(v) - 1], rec(v[:len(v) - 1]))
	}

	return "[" + rec(vet) + "]"
}

// reverse: inverte os elementos do slice
func reverse(vet []int) {
	vet_r := make([]int, len(vet))
	r := 0
	for i := len(vet) - 1; i >= 0; i-- {
		vet_r[r] = vet[i]
		r++
	}
	for i := 0; i < len(vet); i++ {
		vet[i] = vet_r[i]
	}
}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	acc := 0
	for _, v := range vet {
		acc += v
	}
	return acc
}

// mult: produto dos elementos do slice
func mult(vet []int) int {
	acc := 1
	for _, v := range vet {
		acc *= v
	}
	return acc
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {
	if len(vet) == 0 {
		return -1
	}

	i_menor := 0

	for i := 0; i < 1; i++ {
		menor := vet[i]
		for j := 0; j < len(vet); j++ {
			if menor > vet[j] {
				i_menor = j
			}
		}
	}

	return i_menor
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
