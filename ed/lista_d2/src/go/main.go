package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Node struct {
	Value int
	prev *Node
	next *Node
	root *Node
}

type LList struct {
	root *Node
	size int
}

func (n *Node) Next() *Node {
	if n.next == n.root {
		return nil
	}
	
	return n.next
}


func (n *Node) Prev() *Node {
	if n.prev == n.root {
		return nil
	}
	
	return n.prev
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
		result += fmt.Sprintf("%d", curr.Value)  
		if i != l.size - 1 {
			result += ", "
		}
		curr = curr.next
	}
	
	result += "]"
	return result
}

func (l *LList) PushFront(value int) {
	node := &Node{
		Value: value,
		root: l.root,
	}

	first := l.root.next

	node.prev = l.root
	node.next = first

	l.root.next = node
	first.prev = node

	l.size++
}

func (l *LList) PushBack(value int) {
	node := &Node{
		Value: value,
		root: l.root,
	}

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
	l.root.next = l.root
	l.root.prev = l.root
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

func (l *LList) Front() *Node {
	if l.size == 0 {
		return nil
	}

	return l.root.next
}

func (l *LList) Back() *Node {
	if l.size == 0 {
		return nil
	}

	return l.root.prev
}

func (l *LList) Search(value int) *Node {
	curr := l.root.next
	for i := 0; i < l.size; i++ {
		if curr.Value == value {
			return curr
		}
		curr = curr.next
	}
	 
	return nil
}

func (l *LList) Insert(node *Node, newValue int) {
	newNode := &Node{
		Value: newValue,
		root: l.root,
	}

	prev := node.prev

	newNode.prev = prev
	newNode.next = node

	prev.next = newNode
	node.prev = newNode

	l.size++
}

func (l *LList) Remove(node *Node) {
	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev

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
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.Value = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
