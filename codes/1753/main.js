
class MinHeap {
  constructor() {
    this.heap = [];
  }

  length() {
    return this.heap.length
  }

  isEmpty() {
    return this.heap.length === 0
  }

  pop() {
    if (this.heap.length === 1) return this.heap.pop();
    const root = this.heap[0]
    this.heap[0] = this.heap.pop();
    this.bubbleDown();

    return root
  }

  push(v) {
    this.heap.push(v)
    this.bubbleUp();
  }

  bubbleUp() {
    let i = this.heap.length - 1;
    while(i>0) {
      const parent = Math.floor((i - 1) / 2)
      if (this.heap[i][1] > this.heap[parent][1]) break;
      [this.heap[i], this.heap[parent]] = [this.heap[parent], this.heap[i]];
      i = parent;
    }
  }

  bubbleDown() {
    let i = 0;
    const len = this.heap.length;
    while (true) {
      let left = i * 2 + 1
      let right = i * 2 + 2
      let min = i

      if (left < len && this.heap[left][1] < this.heap[min][1]) min = left
      if (right < len && this.heap[right][1] < this.heap[min][1]) min = right
      if (min === i) break;

      [this.heap[i], this.heap[min]] = [this.heap[min], this.heap[i]]
      i = min;
    }
  }
}

const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n")
const [V, E] = input[0].split(" ").map(Number)
const start = Number(input[1])
const graph = Array.from({length: V+1}, ()=> [])

input.slice(2).forEach((line) => {
  const [u, v, w] = line.split(" ").map(Number)
  graph[u].push([v,w])
})

let min = Array.from({length: V + 1}, () => Infinity)
min[start] = 0
let pq = new MinHeap()
pq.push([start, 0]);

while(pq.length()) {
  const [now, curr] = pq.pop();
  if (min[now] < curr) continue;
  
  for (let [next, weight] of graph[now]) {
    if (min[next] > curr + weight) {
      min[next] = curr + weight;
      pq.push([next, curr + weight])
    }
  }
}

console.log(min.slice(1).map((v) => v===Infinity?"INF":v).join("\n"))
