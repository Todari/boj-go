package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	div5 := n / 5
	for {
		remain := n - div5*5
		if remain%3 != 0 {
			div5--
		} else {
			fmt.Fprint(writer, div5+remain/3)
			break
		}
		if div5 < 0 {
			fmt.Fprint(writer, -1)
			break
		}
	}

}
