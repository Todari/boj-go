package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Coord struct {
	y int
	x int
}

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	var N int
	fmt.Fscanln(r, &N)

	Map := make([][]bool, N)
	network := make(map[Coord][]Coord)

	for i := 0; i < N; i++ {
		var input string
		fmt.Fscanln(r, &input)
		for _, v := range strings.Split(input, "") {
			if v == "1" {
				Map[i] = append(Map[i], true)
			} else {
				Map[i] = append(Map[i], false)
			}
		}
	}
	result := make([][]int, N)
	for i := 0; i < N; i++ {
		result[i] = make([]int, N)
	}

	getNetwork(Map, network)
	fmt.Fprintln(w, network)
	makeGroup(network, result)

	fmt.Fprintln(w, network)
	fmt.Fprintln(w, result)
}

func makeGroup(network map[Coord][]Coord, result [][]int) {
	counter := 1
	visited := make(map[Coord]bool)

	keys := make([]Coord, 0, len(network))
	for c := range network {
		keys = append(keys, c)
	}
	start := keys[0]
	visited[start] = true

	// for len(network) != 0 {
	dfs(start, network, visited, result, counter)
	counter += 1
	// }
}

func dfs(start Coord, network map[Coord][]Coord, visited map[Coord]bool, result [][]int, counter int) {
	visited[start] = true
	for _, v := range network[start] {
		if !visited[v] {
			visited[v] = true
			dfs(v, network, visited, result, counter)
			if len(network[start]) != 0 {
				network[start] = network[start][1:]
			}
			result[v.y][v.x] = counter
		}
		delete(network, start)
	}
}

func getNetwork(Map [][]bool, network map[Coord][]Coord) {
	for i, v := range Map {
		for j := range v {
			if Map[i][j] {

				from := Coord{y: i, x: j}
				if i != len(Map)-1 && Map[i+1][j] {
					to := Coord{y: i + 1, x: j}
					network[from] = append(network[from], to)
					network[to] = append(network[to], from)
				}
				if j != len(Map[0])-1 && Map[i][j+1] {
					to := Coord{y: i, x: j + 1}
					network[from] = append(network[from], to)
					network[to] = append(network[to], from)
				}
			}
		}
	}
}
