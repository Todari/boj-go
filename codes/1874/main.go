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

	var n int
	fmt.Fscanln(r, &n)
	arr := make([]int, n)
	s := NewStack()
	for i := 0; i < n; i++ {
		fmt.Fscanln(r, &arr[i])
	}
	var result []string
	counter := 0

	for i := 0; i < n; i++ {
		s.Push(i + 1)
		result = append(result, "+")
		// fmt.Println(s.v.Back().Value, arr[counter])
		for s.v.Back().Value == arr[counter] {
			s.Pop()
			result = append(result, "-")
			counter += 1
			if s.v.Back() == nil {
				break
			}
		}
	}

	if len(result) == 2*n {
		for _, v := range result {
			fmt.Fprintln(w, v)
		}
	} else {
		fmt.Fprintln(w, "NO")
	}
}
