package main

import "fmt"

type Node struct {
	Value string
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Size int
}

func (l *LinkedList) Add(value string) {
	newNode := Node{
		Value: value,
		Next:  l.Head,
	}

	l.Head = &newNode
	l.Size++
}

func (l *LinkedList) DeleteFirst() {
	if l.Head == nil {
		return
	}
	l.Head = l.Head.Next
	l.Size--
}

func (l *LinkedList) Search(value string) *Node {
	current := l.Head
	for current != nil {
		if current.Value == value {
			return current
		}
		current = current.Next
	}
	return nil
}

func (l *LinkedList) List() {
	current := l.Head
	for current != nil {
		fmt.Printf("%+v\n", current)
		current = current.Next
	}
}

func (l *LinkedList) Delete(value string) {
	current := l.Head
	previous := l.Head
	for current != nil {
		if current.Value == value {
			previous.Next = current.Next
			l.Size--
			return
		}
		previous = current
		current = current.Next
	}
}

func main() {
	ll := LinkedList{}

	ll.Add("one")
	ll.Add("two")
	ll.Add("three")

	ll.DeleteFirst()
	ll.List()

	fmt.Println(ll.Search("two"))
}
