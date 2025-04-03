const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n")
const N = Number(input[0])
const map = input.slice(1).map((v) => v.split(" ").map(Number))
const cases = [[0, 1], [0, 1, 2], [1, 2]]

let minDp = [...map[0]]
let maxDp = [...map[0]]

for (let i = 1; i< N; i++) {
  for (let j = 0; j<3; j++)  {
    let next = map[i].filter((_, index) => cases[j].includes(index))

    minDp[j] = minDp[j] + Math.min(...next)
    maxDp[j] = maxDp[j] + Math.max(...next)
  }
}

console.log(Math.max(...maxDp), Math.min(...minDp))

// node.js로 풀 수 없는 문제...?