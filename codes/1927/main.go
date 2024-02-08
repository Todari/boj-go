package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	var N int
	fmt.Fscanln(r, &N)
	mh := &minHeap{}
	heap.Init(mh)

	for i := 0; i < N; i++ {
		var x int
		fmt.Fscanln(r, &x)
		if x != 0 {
			heap.Push(mh, x)
		} else {
			if mh.Len() == 0 {
				fmt.Fprintln(w, 0)
			} else {
				fmt.Fprintln(w, heap.Pop(mh))
			}
		}

	}
}
