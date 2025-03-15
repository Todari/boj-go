package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	var inputs []string
	for {
		// r.ReadString('\n') 은 문자열을 읽어오는 함수
		// 문자열을 읽어오면 문자열 끝에 개행 문자가 포함되어 있음
		// 개행 문자는 문자열에서 제거하기 위해 strings.TrimSpace 함수를 사용
		input, _ := r.ReadString('\n')
		input = strings.TrimSpace(input)
		
		if input == "#" {
			break
		}
		inputs = append(inputs, input)
	}

	for _, input := range inputs {
		count := 0
		for _, vowel := range []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"} {
			count += strings.Count(input, vowel)
		}
		fmt.Fprintln(w, count)
	}
}
