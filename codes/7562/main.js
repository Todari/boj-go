const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().split("\n");
const cases = Number(input[0])
const directions = [[2, 1],[1,2],[2,-1],[1,-2],[-1,-2],[-2,-1],[-2,1],[-1,2]]
const answers = [];

for (let i = 0; i<cases; i++) {
  let N = input[3*i + 1]
  let [sy, sx] = input[3*i + 2].split(" ").map(Number)
  let [ty, tx] = input[3*i + 3].split(" ").map(Number)

  let minimum = Array.from({length: N}, () => Array.from({length: N}).fill(Infinity))
  let visited = Array.from({length: N}, () => Array.from({length: N}).fill(false))
  let queue = [[sy, sx]]
  visited[sy][sx] = true;
  minimum[sy][sx] = 0;

  while (queue.length > 0) {
    let [cy, cx] = queue.shift();

    for (let [dy, dx] of directions) {
      let ny = cy + dy
      let nx = cx + dx

      if (0 <= ny && ny < N && 0 <= nx && nx < N && !visited[ny][nx]) {
        queue.push([ny, nx])
        visited[ny][nx] = true;
        minimum[ny][nx] = minimum[cy][cx] + 1
      }
    }
  }

  console.log(minimum[ty][tx])
}