const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n");
const [R, C] = input[0].split(" ").map(Number)
const map = input.slice(1).map(line => line.split(""))

const directions = [[0, 1], [1, 0], [0, -1], [-1, 0]]

let visitedAlpha = Array(26).fill(false)
visitedAlpha[map[0][0].charCodeAt(0) - 65] = true
let max = 0;

const dfs = (cy, cx, count) => {
  max = Math.max(max, count)
  for (const [dy, dx] of directions) {
    let [ny, nx] = [cy + dy, cx + dx]
    if (0<=ny && ny<R && 0<=nx && nx<C) {
      const nextChar = map[ny][nx].charCodeAt(0) - 65
      if (visitedAlpha[nextChar]) continue
      visitedAlpha[nextChar] = true
      dfs(ny, nx, count+1)
      visitedAlpha[nextChar] = false
    }
  }
}

dfs(0, 0, 1)
console.log(max)