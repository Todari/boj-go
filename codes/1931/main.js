const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n")
const N = Number(input[0])
const times = input.slice(1).map(line => line.split(" ").map(Number))
times.sort((a,b) => {
  if (a[1] === b[1]) {
    return a[0]-b[0]
  }
  return a[1]-b[1]
})

let endTime = 0
let count = 0;
for (const [start, end] of times) {
  if (start >= endTime) {
    endTime = end
    count++;
  }
}

console.log(count)