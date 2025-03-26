const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n")
const N = Number(input[0])
const M = Number(input[1])
const graph = new Map()

class PQ {
  constructor() {
    this.queue = [];
  }

  length() {
    return this.queue.length
  }

  push(v) {
    this.queue.push(v)
    this.bubbleUp();
  }

  pop() {
    if (this.queue.length === 0) return null;
    if (this.queue.length === 1) return this.queue.pop()

    let min = this.queue[0]
    this.queue[0] = this.queue.pop();
    this.bubbleDown();

    return min
  }

  bubbleUp() {
    let index = this.queue.length - 1
    let parentIndex = Math.floor((index - 1) / 2)

    while(index > 0) {
      if (this.queue[index][1] >= this.queue[parentIndex][1]) break;

      [this.queue[index], this.queue[parentIndex]] = [this.queue[parentIndex], this.queue[index]]
      index = parentIndex;
    }
  }

  bubbleDown() {
    let index = 0;
    
    while (true) {
      let left = index * 2 + 1
      let right = index * 2 + 2
      let max = index
      if (left < this.queue.length && this.queue[index][1] > this.queue[left][1]) {
        max = left
      }
      if (right < this.queue.length && this.queue[index][1] > this.queue[right][1]) {
        max = right
      }
      if (max === index) break;

      [this.queue[index], this.queue[max]] = [this.queue[max], this.queue[index]]
      index = max
    }
  }
}

input.slice(2, M+2).map(v => {
  let [start, end, cost] = v.split(" ").map(Number)
  if (!graph.has(start)) graph.set(start, [])
  
  let arr = graph.get(start)
  arr.push([end, cost])
  graph.set(start, arr)
})
const [start, end] = input[M+2].split(" ").map(Number)

let min = Array.from({length: N + 1}, ()=>Infinity)
let queue = new PQ

queue.push([start, 0])
min[start]= 0

while(queue.length() > 0) {
  let [now, cost] = queue.pop();
  if (cost > min[now]) continue;
  if (!graph.has(now)) continue;

  for (let [next, nextCost] of graph.get(now)) {
    if (min[next] > cost + nextCost) {
      min[next] = cost + nextCost
      queue.push([next, cost + nextCost])
    }
  }
}

console.log(min[end])
