package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	Elements []int
}

func (q *Queue) Enqueue(element int) {
	q.Elements = append(q.Elements, element)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.Elements) == 0 {
		return 0, errors.New("queue is empty")
	}
	element := q.Elements[0]
	q.Elements = q.Elements[1:]
	return element, nil
}

func (q *Queue) Lenght() int {
	return len(q.Elements)
}

func (q *Queue) Peek() (int, error) {
	if len(q.Elements) == 0 {
		return 0, errors.New("queue is empty")
	}
	return q.Elements[0], nil
}

func (q *Queue) IsEmpty() bool {
	return q.Lenght() == 0
}

func main() {
	queue := Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println(queue)
	element, _ := queue.Dequeue()
	fmt.Println(element)
	fmt.Println(queue)
}
