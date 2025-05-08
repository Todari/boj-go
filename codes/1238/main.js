const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n")
const [N, M, X] = input[0].split(" ").map(Number)
let graph = Array.from({length : N + 1}, () => [])
let reversedGraph = Array.from({length : N + 1}, () => [])
for (const [start, end, time] of input.slice(1).map(v => v.split(" ").map(Number))) {
  graph[start].push([end, time])
  reversedGraph[end].push([start, time])
}

class MinHeap {
  constructor() {
    this.heap = []
  }

  isEmpty() {
    return !this.heap.length
  }

  pop() {
    if (this.heap.length === 0) return null
    if (this.heap.length === 1) return this.heap.pop()
    let top = this.heap[0];
    let last = this.heap.pop();
    this.heap[0] = last;
    this._bubbleDown();
    return top;
  }

  push(v) {
    this.heap.push(v);
    this._bubbleUp();
  }

  _bubbleDown() {
    let idx = 0;
    while(true) {
      let left = idx * 2 + 1
      let right = idx * 2 + 2
      let min  = idx

      if (right < this.heap.length && this.heap[min][1] > this.heap[right][1]) min = right;
      if (left < this.heap.length && this.heap[min][1] > this.heap[left][1]) min = left;
      if (min === idx) break;

      [this.heap[min], this.heap[idx]] = [this.heap[idx], this.heap[min]]
      idx = min
    }
  }

  _bubbleUp() {
    let idx = this.heap.length - 1
    while(idx > 0) {
      let parent = Math.floor((idx - 1) / 2)

      if (this.heap[parent][1] <= this.heap[idx][1]) break;
      [this.heap[parent], this.heap[idx]] = [this.heap[idx], this.heap[parent]]
      idx = parent;
    }
  } 
}

function dijkstra(start, graph) {
  const dist = Array(N + 1).fill(Infinity);
  const pq = new MinHeap();
  dist[start] = 0;
  pq.push([start, 0]);

  while (!pq.isEmpty()) {
    const [cur, cost] = pq.pop();
    if (dist[cur] < cost) continue;

    for (const [next, nextCost] of graph[cur]) {
      if (dist[next] > cost + nextCost) {
        dist[next] = cost + nextCost;
        pq.push([next, dist[next]]);
      }
    }
  }

  return dist;
}

const distFromX = dijkstra(X, graph);        // X → i
const distToX = dijkstra(X, reversedGraph);  // i → X

let answer = 0;
for (let i = 1; i <= N; i++) {
  answer = Math.max(answer, distFromX[i] + distToX[i]);
}

console.log(answer);