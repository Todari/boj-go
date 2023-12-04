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

func isBalanced(s string) bool {
	stk := NewStack()

	for _, v := range s {
		if v == '(' || v == '[' {
			stk.Push(v)
		} else if v == ')' {
			if stk.v.Back() == nil {
				return false
			} else if stk.v.Back().Value == '(' {
				if stk.Pop() == nil {
					return false
				}
			} else {
				return false
			}
		} else if v == ']' {
			if stk.v.Back() == nil {
				return false
			} else if stk.v.Back().Value == '[' {
				if stk.Pop() == nil {
					return false
				}
			} else {
				return false
			}
		}
	}

	if stk.v.Len() == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var result []bool
	for {
		s.Scan()
		str := s.Text()
		if str == "." {
			break
		}
		result = append(result, isBalanced(str))
	}

	for _, v := range result {
		if v {
			fmt.Fprintln(w, "yes")
		} else {
			fmt.Fprintln(w, "no")
		}
	}
}
