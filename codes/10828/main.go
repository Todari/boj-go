package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	v *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (s *Stack) Push(v interface{}) {
	s.v.PushBack(v)
}

func (s *Stack) Pop() interface{} {
	back := s.v.Back()
	if back == nil {
		return -1
	}

	return s.v.Remove(back)
}

func (s *Stack) Size() interface{} {
	return s.v.Len()
}

func (s *Stack) Empty() interface{} {
	if s.v.Front() == nil {
		return 1
	} else {
		return 0
	}
}

func (s *Stack) Top() interface{} {
	if s.v.Back() == nil {
		return -1
	} else {
		return s.v.Back().Value
	}
}

func main() {
	w := bufio.NewWriter(os.Stdout)
	s := bufio.NewScanner(os.Stdin)
	defer w.Flush()

	s.Scan()
	N, _ := strconv.Atoi(s.Text())

	stk := NewStack()
	for i := 0; i < N; i++ {
		s.Scan()
		input := s.Text()

		switch strings.Split(input, " ")[0] {
		case "push":
			number, _ := strconv.Atoi(strings.Split(input, " ")[1])
			stk.Push(number)
		case "pop":
			fmt.Fprintln(w, stk.Pop())
		case "size":
			fmt.Fprintln(w, stk.Size())
		case "empty":
			fmt.Fprintln(w, stk.Empty())
		case "top":
			fmt.Fprintln(w, stk.Top())
		}
	}
}
