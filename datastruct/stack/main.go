package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	Elements []int
}

func (s *Stack) Push(element int) {
	s.Elements = append(s.Elements, element)
}

func (s *Stack) Pop() (int, error) {
	if len(s.Elements) == 0 {
		return 0, errors.New("stack is empty")
	}
	element := s.Elements[len(s.Elements)-1]
	s.Elements = s.Elements[:len(s.Elements)-1]
	return element, nil
}

func (s *Stack) Peek() (int, error) {
	if len(s.Elements) == 0 {
		return 0, errors.New("stack is empty")
	}
	return s.Elements[len(s.Elements)-1], nil
}

func (s *Stack) Lenght() int {
	return len(s.Elements)
}

func (s *Stack) IsEmpty() bool {
	return len(s.Elements) == 0
}

func main() {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println(stack)
	element, _ := stack.Pop()
	fmt.Println(element)
	fmt.Println(stack)
	element, _ = stack.Pop()
	fmt.Println(element)
	fmt.Println(stack)
	element, _ = stack.Pop()
	fmt.Println(element)
	fmt.Println(stack)
}
