package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Node struct {
	value int
	prev *Node
	next *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	root := &Node{}
	root.next = root
	root.prev = root

	return &LList{
		root: root,
		size: 0,
	}
}

func (l *LList) String() string {
	if l.size == 0 {
		return "[]"
	}

	result := "["
	curr := l.root.next

	for i := 0; i < l.size; i++ {
		result += fmt.Sprintf("%d", curr.value)  
		if i != l.size - 1 {
			result += ", "
		}
		curr = curr.next
	}
	
	result += "]"
	return result
}

func (l *LList) PushFront(value int) {
	node := &Node{value: value}

	first := l.root.next

	node.prev = l.root
	node.next = first

	l.root.next = node
	first.prev = node

	l.size++
}

func (l *LList) PushBack(value int) {
	node := &Node{value: value}

	last := l.root.prev

	node.prev = last
	node.next = l.root

	l.root.prev = node
	last.next = node

	l.size++
}

func (l *LList) Size() int {
	return l.size
}

func (l *LList) Clear() {
	l.size = 0
}

func (l *LList) PopFront() {
	if l.size == 0 {
		return
	}

	first := l.root.next

	l.root.next = first.next
	first.next.prev = l.root

	l.size--
}

func (l *LList) PopBack() {
	if l.size == 0 {
		return
	}

	last := l.root.prev

	last.prev.next = l.root
	l.root.prev = last.prev

	l.size--
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
