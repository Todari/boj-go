package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// isCompleted: 색종이가 모두 같은 색인지 확인
func isCompleted(array []int) bool {
	first := array[0]
	for _, val := range array {
		if val != first {
			return false
		}
	}
	return true
}

// cut: 색종이를 4등분
func cut(array []int, n int) [][]int {
	if isCompleted(array) {
		return [][]int{array}
	}

	result := make([][]int, 4)
	half := n / 2

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			index := i*n + j
			if i < half && j < half {
				result[0] = append(result[0], array[index])
			} else if i < half && j >= half {
				result[1] = append(result[1], array[index])
			} else if i >= half && j < half {
				result[2] = append(result[2], array[index])
			} else {
				result[3] = append(result[3], array[index])
			}
		}
	}
	return result
}

func main() {
	// 입력 받기
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(text))

	colors := make([][]int, 1)
	colors[0] = []int{}

	// 색종이 정보 입력
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		nums := strings.Fields(line)
		for _, num := range nums {
			val, _ := strconv.Atoi(num)
			colors[0] = append(colors[0], val)
		}
	}

	// step 계산
	step := 0
	for temp := n; temp > 1; temp /= 2 {
		step++
	}

	// 색종이 나누기
	for i := 0; i < step; i++ {
		var tmp [][]int
		for _, c := range colors {
			tmp = append(tmp, cut(c, n)...)
		}
		colors = tmp
		n /= 2
	}

	// 결과 저장
	white, blue := 0, 0
	for _, array := range colors {
		if isCompleted(array) {
			if array[0] == 0 {
				white++
			} else {
				blue++
			}
		}
	}

	// 출력
	fmt.Println(white)
	fmt.Println(blue)
}