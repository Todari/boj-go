package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Queue struct {
	v *list.List
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func (q *Queue) Push(v interface{}) {
	q.v.PushBack(v)
}

func (q *Queue) Pop() interface{} {
	front := q.v.Front()
	if front == nil {
		return -1
	}

	return q.v.Remove(front)
}

func (q *Queue) Size() interface{} {
	return q.v.Len()
}

func (q *Queue) Empty() interface{} {
	if q.v.Back() == nil {
		return 1
	} else {
		return 0
	}
}

func (q *Queue) Front() interface{} {
	if q.v.Front() == nil {
		return -1
	} else {
		return q.v.Front().Value
	}
}

func (q *Queue) Back() interface{} {
	if q.v.Back() == nil {
		return -1
	} else {
		return q.v.Back().Value
	}
}

func main() {
	w := bufio.NewWriter(os.Stdout)
	s := bufio.NewScanner(os.Stdin)
	defer w.Flush()

	s.Scan()
	N, _ := strconv.Atoi(s.Text())

	q := NewQueue()
	for i := 0; i < N; i++ {
		s.Scan()
		input := s.Text()

		switch strings.Split(input, " ")[0] {
		case "push":
			number, _ := strconv.Atoi(strings.Split(input, " ")[1])
			q.Push(number)
		case "pop":
			fmt.Fprintln(w, q.Pop())
		case "size":
			fmt.Fprintln(w, q.Size())
		case "empty":
			fmt.Fprintln(w, q.Empty())
		case "front":
			fmt.Fprintln(w, q.Front())
		case "back":
			fmt.Fprintln(w, q.Back())
		}
	}
}
