package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)

	var pq []int
	var mq []int
	var yq []int

	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(r, &x)

		if x > 0 {
			pq = append(pq, x)
			sort.Slice(pq, func(i, j int) bool {
				return pq[i] > pq[j]
			})
		} else if x < 0 {
			mq = append(mq, x)
			sort.Slice(mq, func(i, j int) bool {
				return mq[i] < mq[j]
			})
		} else {
			if len(pq)+len(mq) == 0 {
				yq = append(yq, 0)
			} else if len(pq) == 0 {
				yq = append(yq, mq[len(mq)-1])
				mq = mq[:len(mq)-1]
			} else if len(mq) == 0 {
				yq = append(yq, pq[len(pq)-1])
				pq = pq[:len(pq)-1]
			} else {
				if pq[len(pq)-1] < -mq[len(mq)-1] {
					yq = append(yq, pq[len(pq)-1])
					pq = pq[:len(pq)-1]
				} else {
					yq = append(yq, mq[len(mq)-1])
					mq = mq[:len(mq)-1]
				}
			}
		}
	}

	for _, v := range yq {
		fmt.Fprintln(w, v)
	}
}
