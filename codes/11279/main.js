const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n")
const N = Number(input[0])
let cmds = input.slice(1).map(Number)

class MaxHeap {
  constructor() {
    this.heap = [];
  }

  push(value) {
    this.heap.push(value)
    this._bubbleUp()
  }

  pop() {
    if (this.heap.length === 0) return 0
    if (this.heap.length === 1) return this.heap.pop()

    let max = this.heap[0]
    this.heap[0] = this.heap.pop();
    this._bubbleDown()
    return max;
  }

  _bubbleUp() {
    let index = this.heap.length - 1
    while (index > 0) {
      let parentIndex = Math.floor((index - 1) / 2)
      if (this.heap[index] <= this.heap[parentIndex]) break;

      [this.heap[parentIndex], this.heap[index]] = [this.heap[index], this.heap[parentIndex]]
      index = parentIndex
    }
  }

  _bubbleDown() {
    let index = 0;
    let length = this.heap.length
    while (true){
      let left = index * 2 + 1
      let right = index * 2 + 2
      let max = index;

      if (left < length && this.heap[left] > this.heap[max]) {
        max = left
      }
      if (right < length && this.heap[right] > this.heap[max]) {
        max = right
      }
      if (max === index) break;

      [this.heap[index], this.heap[max]] = [this.heap[max], this.heap[index]]
      index = max
    }
  }
}

let maxHeap = new MaxHeap();
let answer = [];
for (let cmd of cmds) {
  if (cmd === 0) {
    answer.push(maxHeap.pop());
  } else {
    maxHeap.push(cmd)
  }
}

console.log(answer.join("\n"))