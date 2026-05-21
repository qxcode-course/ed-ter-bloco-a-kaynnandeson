package main

import "fmt"

func podeUsar(seq []rune, pos int, num rune, l int) bool {

	inicio := pos - l
	fim := pos + l

	if inicio < 0 {
		inicio = 0
	}

	if fim >= len(seq) {
		fim = len(seq) - 1
	}

	for i := inicio; i <= fim; i++ {

		if i == pos {
			continue
		}

		if seq[i] == num {
			return false
		}
	}

	return true
}

func resolver(seq []rune, l int, pos int) bool {

	// terminou
	if pos == len(seq) {
		return true
	}

	// já preenchido
	if seq[pos] != '.' {
		return resolver(seq, l, pos+1)
	}

	// tenta números de 0 até L
	for d := 0; d <= l; d++ {

		num := rune(d + '0')

		if podeUsar(seq, pos, num, l) {

			seq[pos] = num

			// continua
			if resolver(seq, l, pos+1) {
				return true
			}

			// backtrack
			seq[pos] = '.'
		}
	}

	return false
}

func main() {

	var sequencia string
	var l int

	fmt.Scan(&sequencia)
	fmt.Scan(&l)

	seq := []rune(sequencia)

	resolver(seq, l, 0)

	fmt.Println(string(seq))
}