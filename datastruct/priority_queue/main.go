package main

import (
	"errors"
	"fmt"
)

type PriorityQueue struct {
	High []int
	Low  []int
}

func (q *PriorityQueue) Enqueue(element int, highpriority bool) {
	if highpriority {
		q.High = append(q.High, element)
	} else {
		q.Low = append(q.Low, element)
	}
}

func (q *PriorityQueue) Dequeue() (int, error) {
	if len(q.High) != 0 {
		var firstElement int
		firstElement, q.High = q.High[0], q.High[1:]
		return firstElement, nil
	}
	if len(q.Low) != 0 {
		var firstElement int
		firstElement, q.Low = q.Low[0], q.Low[1:]
		return firstElement, nil
	}

	return 0, errors.New("queue is empty")
}

func (q *PriorityQueue) Lenght() int {
	return len(q.High) + len(q.Low)
}

func (q *PriorityQueue) Peek() (int, error) {
	if len(q.High) != 0 {
		return q.High[0], nil
	}
	if len(q.Low) != 0 {
		return q.Low[0], nil
	}

	return 0, errors.New("queue is empty")
}

func (q *PriorityQueue) IsEmpty() bool {
	return q.Lenght() == 0
}

func main() {
	queue := PriorityQueue{}

	fmt.Println(queue)
	queue.Enqueue(1, true)
	fmt.Println(queue)
	queue.Enqueue(10, false)
	fmt.Println(queue)

	element, _ := queue.Dequeue()
	fmt.Println(element)
	fmt.Println(queue)
}
