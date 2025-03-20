const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().split("\n")
const N = Number(input[0])
const map = input.slice(1).map((line) => line.split(" ").map(Number))
const directions = [[0,1], [1,0], [0,-1], [-1,0]]

let answer = 0;
for (let h = 1; h<=100; h++) {
  let visited = Array.from({length : N}, () => Array.from({length : N}).fill(false))
  let count = 0;
  for (let i = 0; i<N; i++) {
    for (let j = 0; j<N; j++) {
      if (map[i][j] >= h && !visited[i][j]) {
        let queue = [[i, j]]
        visited[i][j] = true;

        while(queue.length > 0) {
          let [cy, cx] = queue.shift()
          for (let [dy, dx] of directions) {
            let y = cy + dy
            let x = cx + dx

            if (0 <= y && y < N && 0 <= x && x < N && !visited[y][x] && map[y][x] >= h) {
              queue.push([y, x])
              visited[y][x] = true;
            }
          }
        }
        
        count++;
      }
    }
  }

  answer = Math.max(answer, count)
}

console.log(answer)