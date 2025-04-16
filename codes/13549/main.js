const [N, K] = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split(" ").map(Number)

let min = new Map();
let queue = [[N, 0]];

while(queue.length > 0) {
  const [pos, time] = queue.shift();
  
  if (pos === K) {
    console.log(time)
    break;
  }
  
  if (min.has(pos) && min.get(pos) <= time) continue
  min.set(pos, time)

  if (pos * 2 <= K + 1 ) {
    queue.unshift([pos * 2, time])
  }
  if (pos + 1 <= K + 1 ) {
    queue.push([pos + 1, time + 1])
  }
  if (pos - 1 >= 0 ) {
    queue.push([pos - 1, time + 1])
  }
}