package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	count := 0
	for _, v := range vet {
		if v > 0 {
			count++
		}
	}

	men := make([]int, count)
	j := 0
	for _, v := range vet {
		if v > 0 {
			men[j] = v
			j++
		}
	}

	return men
}

func getCalmWomen(vet []int) []int {
	count := 0;
	for _, v := range vet {
		if v < 0 && v > -10 {
			count++
		}
	}

	c_women := make([]int, count)
	j := 0
	for _, v := range vet {
		if v < 0 && v > -10 {
			c_women[j] = v
			j++
		}
	}

	return c_women
}

func sortVet(vet []int) []int {
	for i := 0; i < len(vet); i++ {
		for j := i + 1; j < len(vet); j++ {
			if vet[i] > vet[j] {
				aux := vet[i]
				vet[i] = vet[j]
				vet[j] = aux
			} 
		}
	}

	return vet
}

func abs(value int) int {
	if value == 0 {
		return 0
	}
	if value > 0 {
		return value
	}
	return value * -1
}

func sortStress(vet []int) []int {
	for i := 0; i < len(vet); i++ {
		for j := i; j < len(vet); j++ {
			if abs(vet[i]) > abs(vet[j]) {
				aux := vet[i]
				vet[i] = vet[j]
				vet[j] = aux
			}
		}
	}

	return vet
}

func reverse(vet []int) []int {
	_ = vet
	return nil
}

func unique(vet []int) []int {
	_ = vet
	return nil
}

func repeated(vet []int) []int {
	_ = vet
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

