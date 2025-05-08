const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n");
const [N, M] = input[0].split(" ").map((v) => Number(v) - 1)
const map = input.slice(1).map((line) => line.split("").map(Number))

let dist = Array.from({length: N+1}, () => Array(M+1).fill(Infinity))
const directions = [[0, 1], [1, 0], [0, -1], [-1, 0]]
const queue = [[0,0,1]];
dist[0][0] = 1;

while (queue.length > 0) {
  const [y, x, len] = queue.shift();
  for (const [dy, dx] of directions) {
    const [cy, cx] = [y + dy, x + dx]
    if (0 <= cy && cy <= N && 0<=cx && cx <= M && map[cy][cx] !== 0 && len + 1 < dist[cy][cx]) {
      dist[cy][cx] = len + 1
      queue.push([cy, cx, len + 1])
    }
  }
}

console.log(dist[N][M])