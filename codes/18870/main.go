package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// 빠른 입력을 위한 bufio.Reader 생성
	reader := bufio.NewReader(os.Stdin)
	
	// 첫 번째 줄: N
	line, _ := reader.ReadString('\n')
	N, _ := strconv.Atoi(strings.TrimSpace(line))
	
	// 두 번째 줄: 공백으로 구분된 숫자 문자열
	line, _ = reader.ReadString('\n')
	fields := strings.Fields(line)
	arr := make([]int, N)
	for i, s := range fields {
		arr[i], _ = strconv.Atoi(s)
	}
	
	// arr에서 중복 제거한 후 오름차순으로 정렬된 고유 값(unique)을 구함
	uniqueMap := make(map[int]bool)
	for _, v := range arr {
		uniqueMap[v] = true
	}
	uniqueSlice := make([]int, 0, len(uniqueMap))
	for key := range uniqueMap {
		uniqueSlice = append(uniqueSlice, key)
	}
	sort.Ints(uniqueSlice)
	
	// 고유 값에 대해 인덱스를 매핑 (value -> index)
	mapping := make(map[int]int)
	for idx, val := range uniqueSlice {
		mapping[val] = idx
	}
	
	// 원래 arr의 각 원소를 매핑된 인덱스로 변환
	result := make([]string, N)
	for i, v := range arr {
		result[i] = strconv.Itoa(mapping[v])
	}
	
	// 결과 출력
	fmt.Println(strings.Join(result, " "))
}