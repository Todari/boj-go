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

func (q *Stack) Push(v interface{}) {
	q.v.PushBack(v)
}

func (q *Stack) Pop() interface{} {
	back := q.v.Back()
	if back == nil {
		return nil
	}

	return q.v.Remove(back)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var result []string

	var n int
	fmt.Fscan(r, &n)

	for i := 0; i < n; i++ {
		var str string
		stk := NewStack()
		fmt.Fscan(r, &str)

		for i, v := range str {
			if string(v) == "(" {
				stk.Push(string(v))
			} else {
				if stk.Pop() == nil {
					result = append(result, "NO")
					break
				}
			}

			if i == len(str)-1 {
				if stk.v.Len() == 0 {
					result = append(result, "YES")
				} else {
					result = append(result, "NO")
				}
			}
		}
	}

	for _, v := range result {
		fmt.Fprintln(w, v)
	}
}
