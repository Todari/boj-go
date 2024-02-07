package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
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
		return nil
	}

	return s.v.Remove(back)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var K int
	fmt.Fscanln(r, &K)

	stk := NewStack()
	for i := 0; i < K; i++ {
		var N int
		fmt.Fscanln(r, &N)

		if N == 0 {
			stk.Pop()
		} else {
			stk.Push(N)
		}
	}

	result := 0
	for p := stk.v.Front(); p != nil; p = p.Next() {
		result += p.Value.(int)
	}

	fmt.Fprintln(w, result)
}
