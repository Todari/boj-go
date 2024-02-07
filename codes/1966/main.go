package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

type Docs struct {
	idx        int
	importance int
}

func NewPrintQueue(s []int) *PrintQueue {
	li := list.New()
	for i, v := range s {
		li.PushBack(Docs{i, v})
	}
	return &PrintQueue{li}
}

func (pq *PrintQueue) getProcedure(idx int) int {
	counter := 0
	for pq.v.Front() != nil {
		maxImportance := 0
		for p := pq.v.Front(); p != nil; p = p.Next() {
			if p.Value.(Docs).importance > maxImportance {
				maxImportance = p.Value.(Docs).importance
			}
		}
		for p := pq.v.Front(); p != nil; p = p.Next() {
			front := pq.v.Remove(pq.v.Front())
			if front.(Docs).importance != maxImportance {
				pq.v.PushBack(front)
			} else {
				counter += 1
				if front.(Docs).idx == idx {
					return counter
				}
			}
		}
	}
	return -1
}

type PrintQueue struct {
	v *list.List
}

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	var T int
	fmt.Fscanln(r, &T)

	for i := 0; i < T; i++ {
		var N, M int
		fmt.Fscan(r, &N)
		fmt.Fscan(r, &M)
		fmt.Fscanln(r)
		arr := make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Fscan(r, &arr[j])
		}
		fmt.Fscanln(r)
		pq := NewPrintQueue(arr)
		fmt.Fprintln(w, pq.getProcedure(M))
	}
}
