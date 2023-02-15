package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	top *Node
}

func (s *Stack) Push(value int) {
	node := &Node{value, s.top}
	s.top = node
}

func (s *Stack) Pop() (int, bool) {
	if s.top == nil {
		return 0, false
	}
	value := s.top.value
	s.top = s.top.next
	return value, true
}

func (s *Stack) Peek() (int, bool) {
	if s.top == nil {
		return 0, false
	}
	return s.top.value, true
}

func (s *Stack) Clear() {
	s.top = nil
}

func (s *Stack) Contains(value int) bool {
	current := s.top
	for current != nil {
		if current.value == value {
			return true
		}
		current = current.next
	}
	return false
}

func (s *Stack) Increment() {
	current := s.top
	for current != nil {
		current.value++
		current = current.next
	}
}

func (s *Stack) Print() {
	current := s.top
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}

func (s *Stack) PrintReverse() {
	values := []int{}
	current := s.top
	for current != nil {
		values = append(values, current.value)
		current = current.next
	}
	for i := len(values) - 1; i >= 0; i-- {
		fmt.Println(values[i])
	}
}

func main() {
	s := &Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println("Printing stack:")
	s.Print()
	fmt.Println("Printing stack in reverse:")
	s.PrintReverse()
	fmt.Println("Incrementing values:")
	s.Increment()
	fmt.Println("Printing stack after incrementing values:")
	s.Print()
	fmt.Println("Popping elements:")
	for i := 0; i < 3; i++ {
		value, ok := s.Pop()
		if ok {
			fmt.Println(value)
		}
	}
	fmt.Println("Stack cleared, printing:")
	s.Clear()
	s.Print()
}
