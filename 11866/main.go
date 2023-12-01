package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	var N, K int
	var result []string
	fmt.Fscanf(r, "%d %d", &N, &K)

	li := new(list.List)
	for i := 1; i < N+1; i++ {
		li.PushBack(i)
	}

	p := li.Front()
	for li.Len() > 0 {
		for i := 1; i < K+1; i++ {
			if p.Next() == nil {
				p = li.Front()
			} else {
				p = p.Next()
			}
		}

		if p.Prev() != nil {
			result = append(result, strconv.Itoa(li.Remove(p.Prev()).(int)))
		} else {
			result = append(result, strconv.Itoa(li.Remove(li.Back()).(int)))
		}
	}

	fmt.Fprintln(w, "<"+strings.Join(result, ", ")+">")
}
