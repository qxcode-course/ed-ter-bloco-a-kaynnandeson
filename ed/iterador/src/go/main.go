package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MyList struct {
	data []int
}

type Iterator struct {
	data  []int
	index int
	reverse bool
}

func NewMyList(values []int) *MyList {
	return &MyList{data: values}
}

func (l *MyList) Iterator() *Iterator {
	return &Iterator{data: l.data, index: -1}
}

func (i *Iterator) HasNext() bool {
	if i.reverse {
		return i.index > 0
	}
	return i.index < len(i.data)-1
}

func (i *Iterator) Next() int {
	// if i.index == len(i.data) {
	// 	panic(fmt.Errorf("No more elements"))
	// }

	if len(i.data) == 0 {
		panic("empty list")
	}
	
	if i.reverse {
		i.index--
		return i.data[i.index]
	}

	i.index += 1

	if i.index == len(i.data) {
		i.index = 0
	}

	return i.data[i.index]
}

func (l *MyList) ReverseIterator() *Iterator {
	return &Iterator {
		data: l.data,
		index: len(l.data),
		reverse: true,
	}
}

func (l *MyList) CyclicIterator() *Iterator {
	return &Iterator{
		data: l.data,
		index: -1,
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	mylist := NewMyList([]int{})
	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			for i := 1; i < len(args); i++ {
				slice := make([]int, len(args)-1)
				for i, value := range args[1:] {
					slice[i], _ = strconv.Atoi(value)
				}
				mylist = NewMyList(slice)
			}
		case "show":
			fmt.Print("[ ")
			for it := mylist.Iterator(); it.HasNext(); {
				fmt.Printf("%v ", it.Next())
			}
			fmt.Println("]")
		case "reverse":
			fmt.Print("[ ")
			for it := mylist.ReverseIterator(); it.HasNext(); {
				fmt.Printf("%v ", it.Next())
			}
			fmt.Println("]")
		case "cyclic":
			qtd, _ := strconv.Atoi(args[1])
			fmt.Print("[ ")
			it := mylist.CyclicIterator()
			for range qtd {
				fmt.Printf("%v ", it.Next())
			}
			fmt.Println("]")
		}
	}

}
