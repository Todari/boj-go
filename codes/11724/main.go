package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var	graph map[int][]int
var visited []bool

func main() {
	r:=bufio.NewReader(os.Stdin)
	w:=bufio.NewWriter(os.Stdout)
	defer w.Flush()

	line, _ := r.ReadString('\n')
	parts := strings.Fields(line)
	N, _ := strconv.Atoi(parts[0])
	M, _ := strconv.Atoi(parts[1])

	graph = make(map[int][]int, N + 1)
	visited = make([]bool, N + 1)

	for i:=0; i<M; i++ {
		line, _ := r.ReadString('\n')
		parts := strings.Fields(line)
		u, _ := strconv.Atoi(parts[0])
		v, _ := strconv.Atoi(parts[1])

		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	answer := 0;
	for i := 1; i<= N; i++ {
		if !visited[i] {
			dfs(i)
			answer++
		}
	}

	fmt.Fprintln(w, answer)
}

func dfs(node int) {
	visited[node] = true
	for _, next := range graph[node] {
		if !visited[next] {
			dfs(next)
		}
	}
}
