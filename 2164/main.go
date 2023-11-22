package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
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
		return nil
	}

	return q.v.Remove(front)
}

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)

	q := NewQueue()
	for i := 0; i < n; i++ {
		q.Push(i + 1)
	}

	for q.v.Len() > 1 {
		q.Pop()
		q.Push(q.Pop())
	}
	result := q.Pop()

	fmt.Fprint(w, result)
}
