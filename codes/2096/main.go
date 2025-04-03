package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func max(nums ...int) int {
	result := nums[0]
	for _, v := range nums {
		if v > result {
			result = v
		}
	}
	return result
}

func min(nums ...int) int {
	result := nums[0]
	for _, v := range nums {
		if v < result {
			result = v
		}
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	minDp := make([]int, 3)
	maxDp := make([]int, 3)

	for i := 0; i < N; i++ {
		scanner.Scan()
		line := strings.Split(scanner.Text(), " ")
		nums := make([]int, 3)
		for j := 0; j < 3; j++ {
			nums[j], _ = strconv.Atoi(line[j])
		}

		if i == 0 {
			copy(minDp, nums)
			copy(maxDp, nums)
		} else {
			prevMin := make([]int, 3)
			prevMax := make([]int, 3)
			copy(prevMin, minDp)
			copy(prevMax, maxDp)

			minDp[0] = nums[0] + min(prevMin[0], prevMin[1])
			minDp[1] = nums[1] + min(prevMin[0], prevMin[1], prevMin[2])
			minDp[2] = nums[2] + min(prevMin[1], prevMin[2])

			maxDp[0] = nums[0] + max(prevMax[0], prevMax[1])
			maxDp[1] = nums[1] + max(prevMax[0], prevMax[1], prevMax[2])
			maxDp[2] = nums[2] + max(prevMax[1], prevMax[2])
		}
	}

	fmt.Printf("%d %d\n", max(maxDp...), min(minDp...))
}