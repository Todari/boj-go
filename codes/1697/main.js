const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString()
const [N, K] = input.split(" ").map(Number)
let min = new Map()
min.set(N, 0)

let queue = [N]
while(queue.length > 0) {
  let n = queue.shift()
  let nextPosition = [n-1, n+1, n*2]
  
  for (const next of nextPosition) {
    if (!min.has(next) && next >= 0 && next <= 100000 ) {
      min.set(next, min.get(n) + 1)
      queue.push(next)
    }
  }
}

console.log(min.get(K))