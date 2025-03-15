package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	paper [][]int
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	var N int
	fmt.Fscanln(r, &N)

	paper = make([][]int, 1)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			var x int
			fmt.Fscan(r, &x)
			paper[0] = append(paper[0], x)
		}
		fmt.Fscanln(r)
	}

	// 기존에 주어진 paper가 자를 필요 없어도 자름
	// 수정이 필요해!!!
	cut()
	white, blue := get()
	fmt.Fprintln(w, white)
	fmt.Fprintln(w, blue)
}

func cut() {
START:
	for i, v := range paper {
		if !check(v) {
			result := make([][]int, 4)

			for ii, vv := range v {
				if (ii/(int(math.Sqrt(float64(len(v))))/2))%2 == 0 && ii/(len(v)/2) < 1 {
					result[0] = append(result[0], vv)
				}
				if (ii/(int(math.Sqrt(float64(len(v))))/2))%2 != 0 && ii/(len(v)/2) < 1 {
					result[1] = append(result[1], vv)
				}
				if (ii/(int(math.Sqrt(float64(len(v))))/2))%2 == 0 && ii/(len(v)/2) >= 1 {
					result[2] = append(result[2], vv)
				}
				if (ii/(int(math.Sqrt(float64(len(v))))/2))%2 != 0 && ii/(len(v)/2) >= 1 {
					result[3] = append(result[3], vv)
				}
			}
			paper = append(paper[:i], paper[i+1:]...)
			paper = append(paper, result...)

		}
		fmt.Println(paper)
	}

	for _, v := range paper {
		if !check(v) {
			goto START
		}
	}
}

func check(paper []int) bool {
	init := paper[0]
	for _, v := range paper {
		if v != init {
			return false
		}
	}
	return true
}

func get() (int, int) {
	white, blue := 0, 0
	for _, v := range paper {
		if v[0] == 1 {
			blue += 1
		} else {
			white += 1
		}
	}

	return white, blue
}
