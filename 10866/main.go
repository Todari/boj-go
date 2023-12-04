package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Deque struct {
	v *list.List
}

func NewDeque() *Deque {
	return &Deque{list.New()}
}

func (d *Deque) PushFront(v interface{}) {
	d.v.PushFront(v)
}

func (d *Deque) PushBack(v interface{}) {
	d.v.PushBack(v)
}

func (d *Deque) PopFront() interface{} {
	front := d.v.Front()
	if front == nil {
		return -1
	}

	return d.v.Remove(front)
}

func (d *Deque) PopBack() interface{} {
	back := d.v.Back()
	if back == nil {
		return -1
	}

	return d.v.Remove(back)
}

func (d *Deque) Size() interface{} {
	return d.v.Len()
}

func (d *Deque) Empty() interface{} {
	if d.v.Front() == nil {
		return 1
	}

	return 0
}

func (d *Deque) Front() interface{} {
	if d.v.Front() == nil {
		return -1
	}

	return d.v.Front().Value
}

func (d *Deque) Back() interface{} {
	if d.v.Back() == nil {
		return -1
	}

	return d.v.Back().Value
}

func main() {
	w := bufio.NewWriter(os.Stdout)
	s := bufio.NewScanner(os.Stdin)
	defer w.Flush()

	s.Scan()
	N, _ := strconv.Atoi(s.Text())

	dq := NewDeque()
	for i := 0; i < N; i++ {
		s.Scan()
		input := s.Text()

		switch strings.Split(input, " ")[0] {
		case "push_front":
			number, _ := strconv.Atoi(strings.Split(input, " ")[1])
			dq.PushFront(number)
		case "push_back":
			number, _ := strconv.Atoi(strings.Split(input, " ")[1])
			dq.PushBack(number)
		case "pop_front":
			fmt.Fprintln(w, dq.PopFront())
		case "pop_back":
			fmt.Fprintln(w, dq.PopBack())
		case "size":
			fmt.Fprintln(w, dq.Size())
		case "empty":
			fmt.Fprintln(w, dq.Empty())
		case "front":
			fmt.Fprintln(w, dq.Front())
		case "back":
			fmt.Fprintln(w, dq.Back())
		}
	}
}
