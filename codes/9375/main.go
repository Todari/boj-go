package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var T int
	fmt.Fscanln(r, &T)

	for i := 0; i < T; i++ {
		var n int
		fmt.Fscanln(r, &n)
		typeIndex := make(map[string]int)
		var clothingCount []int

		for j := 0; j < n; j++ {
			var clothes, types string
			fmt.Fscan(r, &clothes)
			fmt.Fscan(r, &types)
			fmt.Fscanln(r)
			idx, ok := typeIndex[types]
			if ok {
				clothingCount[idx]++
			} else {
				typeIndex[types] = len(clothingCount)
				clothingCount = append(clothingCount, 1)
			}
		}

		result := 1
		for _, count := range clothingCount {
			result *= count + 1 // 각 옷 종류별로 선택할 수 있는 옷의 수에 1을 더하여 곱함
		}

		result-- // 모든 종류에서 아무것도 선택하지 않은 경우를 뺌
		fmt.Fprintln(w, result)
	}
}
