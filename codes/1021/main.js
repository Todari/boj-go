const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().split("\n")
const [N, M] = input[0].split(" ").map(Number)
let targets = input[1].split(" ").map(Number)
let dequeue = Array.from({length: N}, (_, k) => k + 1)
let answer = 0;

while (targets.length > 0) {
  let target = targets.shift()
  let idx = dequeue.indexOf(target)
  if (idx < Math.floor((dequeue.length + 1) / 2)) {
    for (let i = 0; i< idx; i++) {
      dequeue.push(dequeue.shift())
      answer++;
    }
  } else {
    for (let i = dequeue.length - 1; i>= idx; i--) {
      dequeue.unshift(dequeue.pop())
      answer++;
    }
  }
  dequeue.shift();
}

console.log(answer)

