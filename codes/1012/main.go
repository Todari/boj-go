package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	graph   [][]int
	visited [][]bool
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	var T int
	fmt.Fscanln(r, &T)

	for j := 0; j < T; j++ {
		counter := 0
		var M, N, K int
		fmt.Fscanf(r, "%d %d %d\n", &M, &N, &K)
		graph = make([][]int, N)
		visited = make([][]bool, N)
		for i := 0; i < N; i++ {
			graph[i] = make([]int, M)
			visited[i] = make([]bool, M)
		}

		for i := 0; i < K; i++ {
			var y, x int
			fmt.Fscan(r, &x)
			fmt.Fscan(r, &y)
			fmt.Fscanln(r)
			graph[y][x] = 1
		}
		//fmt.Fprintln(w, graph)
		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				if graph[i][j] == 1 {
					if !visited[i][j] {
						dfs(i, j)
						counter += 1
					}
				}
			}
		}
		fmt.Fprintln(w, counter)
	}
}

func dfs(y, x int) {
	visited[y][x] = true
	if y != len(graph)-1 && !visited[y+1][x] && graph[y+1][x] == 1 {
		dfs(y+1, x)
	}
	if y != 0 && !visited[y-1][x] && graph[y-1][x] == 1 {
		dfs(y-1, x)
	}
	if x != len(graph[0])-1 && !visited[y][x+1] && graph[y][x+1] == 1 {
		dfs(y, x+1)
	}
	if x != 0 && !visited[y][x-1] && graph[y][x-1] == 1 {
		dfs(y, x-1)
	}
}
