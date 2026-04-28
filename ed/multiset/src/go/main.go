package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Multiset struct {
	data []int
	size int
	capacity int
}

func NewMultiSet(capacity int) *Multiset {
	return &Multiset{
		data: make([]int, capacity),
		size: 0,
		capacity: capacity,
	}
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func (ms *Multiset) String() string {
	if ms.size == 0 {
		return "[]"
	}
	return "[" + Join(ms.data[:ms.size], ", ") + "]" 
}


func (ms *Multiset) lowerBound(value int) int {
	left, right := 0, ms.size

	for left < right {
		mid := (left + right) / 2

		if ms.data[mid] < value {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func (ms *Multiset) Insert(value int) {
	pos := ms.lowerBound(value)

	if ms.size == ms.capacity {
		newCap := 1
		if ms.capacity > 0 {
			newCap = ms.capacity * 2
		}

		newData := make([]int, newCap)
		for i := 0; i < ms.size; i++ {
			newData[i] = ms.data[i]
		}
		ms.data = newData
		ms.capacity = newCap
	}

	for i := ms.size; i > pos; i-- {
		ms.data[i] = ms.data[i - 1]
	}

	ms.data[pos] = value
	ms.size++
}

func (ms *Multiset) Contains(value int) bool {
	for i := 0; i < ms.size; i++ {
		if ms.data[i] == value {
			return true
		}
	}

	return false
}

func (ms *Multiset) Erase(value int) error {
	for i := 0; i < ms.size; i++ {
		if value == ms.data[i] {
			for j := i; j < ms.size - 1; j++ {
				if ms.data[j] == ms.data[j + 1] {
					i++
					continue
				} 
				ms.data[j] = ms.data[i + 1]
			}
			ms.size--
			return nil
		}
	}

	return fmt.Errorf("value not found")
}

func (ms *Multiset) Count(value int) int {
	count := 0
	for i := 0; i < ms.size; i++ {
		if value == ms.data[i] {
			count++
		}
	}

	return count
}

func (ms *Multiset) Clear() {
	ms.size = 0
}


func (ms *Multiset) Unique() int {
	if ms.size == 0 {
		return 0
	}
	
	count := 1

	for i := 1; i < ms.size; i++ {
		if ms.data[i] != ms.data[i - 1] {
			count++
		}
	}

	return count
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	
	ms := NewMultiSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms)
		case "erase":
			value, _ := strconv.Atoi(args[1])
			err := ms.Erase(value)
			if err != nil {
				fmt.Println(err)
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			if ms.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "count":
			value, _ := strconv.Atoi(args[1])
			qtd := ms.Count(value)
			fmt.Println(qtd)
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
