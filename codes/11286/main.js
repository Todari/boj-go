const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n")
const N = Number(input[0])
let inputs = input.slice(1).map(Number)

class AbsHeap {
  constructor() {
    this.heap = [];
  }
  
  push(value) {
    this.heap.push(value);
    this.bubbleUp();
  }
  
  pop() {
    if (this.heap.length === 0) return 0;
    if (this.heap.length === 1) return this.heap.pop();
    
    const min = this.heap[0];
    this.heap[0] = this.heap.pop();
    this.bubbleDown();
    return min;
  }
  
  bubbleUp() {
    let index = this.heap.length - 1;
    while (index > 0) {
      const parentIndex = Math.floor((index - 1) / 2);
      // swap
      if (
        Math.abs(this.heap[parentIndex]) < Math.abs(this.heap[index]) ||
        (Math.abs(this.heap[parentIndex]) === Math.abs(this.heap[index]) && this.heap[parentIndex] < this.heap[index])
      ) break;

      [this.heap[parentIndex], this.heap[index]] = [this.heap[index], this.heap[parentIndex]];
      index = parentIndex;
    }
  }
  
  // 아래로 내리는 과정: 부모 노드와 자식 노드들을 비교하여 더 작은 자식과 교환하며 힙 순서를 유지
  bubbleDown() {
    let index = 0;
    const length = this.heap.length;
    while (true) {
      let left = 2 * index + 1;
      let right = 2 * index + 2;
      let smallest = index;
      
      if (left < length &&
        (Math.abs(this.heap[left]) < Math.abs(this.heap[smallest]) ||
        (Math.abs(this.heap[left]) === Math.abs(this.heap[smallest]) && this.heap[left] < this.heap[smallest] ))
      ) {
        smallest = left;
      }
      if (right < length &&
        (Math.abs(this.heap[right]) < Math.abs(this.heap[smallest]) ||
        (Math.abs(this.heap[right]) === Math.abs(this.heap[smallest]) && this.heap[right] < this.heap[smallest] ))
      ) {
        smallest = right;
      }
      if (smallest === index) break;
      
      // swap
      [this.heap[index], this.heap[smallest]] = [this.heap[smallest], this.heap[index]];
      index = smallest;
    }
  }
}

let heap = new AbsHeap();
let result = [];

for (let char of inputs) {
  // let popped = 0
  if (char !== 0) {
    heap.push(char)
  } else {
    result.push(heap.pop())
  }
}

console.log(result.join("\n"))