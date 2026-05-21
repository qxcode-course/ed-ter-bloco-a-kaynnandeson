package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	One int
	Two int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ordenar(vet []Pair) {
	for i := 0; i < len(vet); i++ {
		for j := i + 1; j < len(vet); j++ {
			if vet[j].One < vet[i].One {
				vet[i], vet[j] = vet[j], vet[i]
			}			
		}
	}
}

func occurr(vet []int) []Pair {
	var pVet []Pair

	for i := 0; i < len(vet); i++ {
		var count Pair
		if vet[i] == 0 {
			continue
		}
		count.One = abs(vet[i])
		count.Two = 1
		for j := i + 1; j < len(vet); j++ {
			if count.One == abs(vet[j]) {
				count.Two++
				vet[j] = 0
			}
		}
		pVet = append(pVet, count)
	}

	ordenar(pVet)

	return pVet
}

func teams(vet []int) []Pair {
	var pVet []Pair

	for i := 0; i < len(vet); i++ {
		var count Pair
		if vet[i] == 0 {
			continue
		}
		count.One = abs(vet[i])
		count.Two = 1
		for j := i + 1; j < len(vet); j++ {
			if vet[j] != vet[i] {
				break	
			}
			if count.One == abs(vet[j]) {
				count.Two++
				vet[j] = 0
			}
		}
		pVet = append(pVet, count)
	}

	return pVet
}

func mnext(vet []int) []int {
	mVet := make([]int, len(vet))
	
	for i := 0; i < len(vet); i++ {
		if vet[i] > 0 {
			if (i > 0 && vet[i - 1] < 0) || (i < len(vet) - 1 && vet[i + 1] < 0) {
				mVet[i] = 1
			}
		}
	}

	return mVet
}

func alone(vet []int) []int {
	aVet := make([]int, len(vet))

	for i := 0; i < len(vet); i++ {
		if vet[i] > 0 {
			if !((i > 0 && vet[i - 1] < 0) || (i < len(vet) - 1 && vet[i + 1] < 0)) {
				aVet[i] = 1
			}
		}
	}

	return aVet
}

func couple(vet []int) int {
	if len(vet) >= 1 {
		return 0
	}

	count := 0

	for i := 0; i < len(vet); i++ {
		if i > 0 {
			if vet[i] == abs(vet[i + 1]) {
				count++
			}
		}
	}

	return count
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	_ = vet
	_ = seq
	_ = pos
	return false
}

func subseq(vet []int, seq []int) int {
	_ = vet
	_ = seq
	return -1
}

func erase(vet []int, posList []int) []int {
	_ = vet
	_ = posList
	return nil
}

func clear(vet []int, value int) []int {
	_ = vet
	_ = value
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
