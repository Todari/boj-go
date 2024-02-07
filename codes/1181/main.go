package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func deleteDuplicated(s []string) []string {
	keys := make(map[string]struct{})
	res := make([]string, 0)
	for _, v := range s {
		if _, ok := keys[v]; ok {
			continue
		} else {
			keys[v] = struct{}{}
			res = append(res, v)
		}
	}
	return res
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N int
	var strArr []string
	fmt.Fscanln(r, &N)

	for i := 0; i < N; i++ {
		var S string
		fmt.Fscanln(r, &S)
		strArr = append(strArr, S)
	}

	strArr = deleteDuplicated(strArr)

	sort.Slice(strArr, func(i, j int) bool {
		if len(strArr[i]) != len(strArr[j]) {
			return len(strArr[i]) < len(strArr[j])
		} else {
			for n := 0; n < len(strArr[i]); n++ {
				if int(strArr[i][n]) < int(strArr[j][n]) {
					return strArr[i] < strArr[j]
				}
			}
		}
		return len(strArr[i]) < len(strArr[j])
	})

	for _, v := range strArr {
		fmt.Fprintln(w, v)
	}
}
