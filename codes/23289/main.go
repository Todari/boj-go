package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {

	s := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	s.Scan()
	R, _ := strconv.Atoi(strings.Split(s.Text(), " ")[0])
	C, _ := strconv.Atoi(strings.Split(s.Text(), " ")[1])
	targetTemp, _ := strconv.Atoi(strings.Split(s.Text(), " ")[2])

	var room [][]int 
	var pan [][]int //y, x, direction
	var targetCoord [][]int //y, x
	var wall [][]int //
	for i := 0; i < R; i++ {
		s.Scan()
		var arr []int
		for _, v := range strings.Split(s.Text(), " ") {
			vInt, _ := strconv.Atoi(v)
			arr = append(arr, vInt)
		}
		room = append(room, arr)
	}
}
