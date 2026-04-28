package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Set struct {
	data []int
	size int 
	capacity int
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}

func (s *Set) lowerBound(value int) int {
	left, right := 0, s.size

	for left < right {
		mid := (left + right) / 2

		if s.data[mid] < value {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func NewSet(capacity int) *Set {
	return &Set{
		data: make([]int, capacity),
		size: 0,
		capacity: capacity,
	}
}

func (s *Set) String() string {
	if s.size == 0 {
		return "[]"
	}
	return "[" + Join(s.data[:s.size], ", ") + "]" 
}

func (s *Set) Insert(value int) {
	pos := s.lowerBound(value)

	if pos < s.size && s.data[pos] == value {
		return
	}

	if s.size == s.capacity {
		newCap := 1
		if s.capacity > 0 {
			newCap = s.capacity * 2
		}

		newData := make([]int, newCap)
		for i := 0; i < s.size; i++ {
			newData[i] = s.data[i]
		}
		s.data = newData
		s.capacity = newCap
	}

	for i := s.size; i > pos; i-- {
		s.data[i] = s.data[i - 1]
	}

	s.data[pos] = value
	s.size++
}

func (s *Set) Contains(value int) bool {
	for i := 0; i < s.size; i++ {
		if s.data[i] == value {
			return true
		}
	}

	return false
}

func (s *Set) Erase(value int) error {
	sucess := false
	for i := 0; i < s.size; i++ {
		if value == s.data[i] {
			for j := i; j < s.size - 1; j++ {	
				s.data[j] = s.data[j + 1]
			}
			s.size--
			sucess = true
		}
	}

	if sucess {
		return nil
	}
	return fmt.Errorf("value not found")
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			fmt.Println(v)
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			err := v.Erase(value)
			if err != nil {
				fmt.Println(err)
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
