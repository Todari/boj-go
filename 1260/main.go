package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"sort"
)

type Graph struct {
	network map[int][]int
	visited map[int]bool
	order   []int
}

func NewGraph() *Graph {
	g := &Graph{}
	g.network = make(map[int][]int)
	g.visited = make(map[int]bool)
	return g
}

func (g *Graph) sortNetwork() {
	for _, v := range g.network {
		sort.Ints(v)
	}
}

func (g *Graph) reset() {
	g.visited = make(map[int]bool)
	g.order = []int{}
}

func (g *Graph) addEdge(from, to int) {
	g.network[from] = append(g.network[from], to)
	g.network[to] = append(g.network[to], from)
}

func (g *Graph) dfs(x int) {
	g.order = append(g.order, x)
	g.visited[x] = true

	for _, v := range g.network[x] {
		if !g.visited[v] {
			g.dfs(v)
		}
	}
}

func (g *Graph) bfs(x int) {
	g.visited[x] = true

	q := new(list.List)
	q.PushBack(x)

	for q.Len() != 0 {
		current := q.Front().Value.(int)
		q.Remove(q.Front())
		g.order = append(g.order, current)

		for _, v := range g.network[current] {
			if !g.visited[v] {
				g.visited[v] = true
				q.PushBack(v)
			}
		}
	}
}

func (g *Graph) print(w *bufio.Writer) {
	for i, v := range g.order {
		if i == len(g.order)-1 {
			fmt.Fprintf(w, "%d\n", v)
			break
		}
		fmt.Fprintf(w, "%d ", v)
	}
}

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	var N, M, V int
	fmt.Fscanf(r, "%d %d %d\n", &N, &M, &V)

	graph := NewGraph()

	for i := 0; i < M; i++ {
		var node1, node2 int
		fmt.Fscan(r, &node1)
		fmt.Fscan(r, &node2)
		fmt.Fscanln(r)

		graph.addEdge(node1, node2)
	}

	graph.sortNetwork()

	graph.dfs(V)
	graph.print(w)
	graph.reset()
	graph.bfs(V)
	graph.print(w)
}
