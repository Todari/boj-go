const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().split("\n")
const [M, N, K] = input[0].split(" ").map(Number);

const squares = input.slice(1).map((line) => line.split(" ").map(Number))

let map = Array.from({length: M}, () => Array.from({length: N}).fill(0))
for (let [x1, y1, x2, y2] of squares) {
  for (let x=x1; x<x2; x++) {
    for (let y=y1; y<y2; y++) {
      map[y][x] = 1
    }
  }
}

let visited = Array.from({length: M}, () => Array.from({length: N}).fill(false))
const directions = [[0, 1], [1, 0], [0, -1], [-1, 0]]
let answer = [];
for (let y = 0; y< M; y++) {
  for (let x = 0; x< N; x++) {
    if (map[y][x] === 0 && !visited[y][x]) {
      let size = 0;
      let queue = [[y,x]]
      visited[y][x] = true

      while(queue.length > 0) {
        let [cy, cx] = queue.shift()
        size++;
        for (let [dy, dx] of directions) {
          let ny = cy + dy
          let nx = cx + dx
          if (0 <= ny && ny < M && 0 <= nx && nx < N && !visited[ny][nx] && map[ny][nx] === 0) {
            visited[ny][nx] = true;
            queue.push([ny, nx])
          }
        }
      }
      answer.push(size)
    }
  }  
}

console.log(answer.length)
console.log(answer.sort((a,b) => a-b).join("\n"))

