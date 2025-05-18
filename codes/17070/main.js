const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n");
const N = Number(input[0])
const arr = input.slice(1).map(line => line.split(" ").map(Number))

let dp = Array.from({length: N},()=> Array.from({length: N}, () => Array.from({length:3}, () => 0)))

// 0 : 가로 1 : 세로 2 : 대각선
// const directions = [[[0,1], [1,1]], [[1,0], [1,1]], [[0,1], [1,1], [1,0]]]
// let queue = [[0, 1, 0]]

// while(queue.length>0) {
//   let [cy, cx, dir] = queue.shift()
//   for (const [dy, dx] of directions[dir]) {
//     const newDir = dx === 1? dy === 1? 2 : 0 : 1
//     let [ny, nx] = [cy+dy, cx+dx]
//     if (0<=ny && ny<N && 0<=nx && nx<N && arr[ny][nx] === 0) {
//       if (newDir === 2 && (arr[ny - 1][nx] !== 0 || arr[ny][nx - 1] !== 0)) continue
//       dp[ny][nx][newDir] += 1
//       queue.push([ny, nx, newDir])
//     }
//   }
// }
dp[0][1][0] = 1;

for (let y = 0; y < N; y++) {
  for (let x = 0; x < N; x++) {
    if (arr[y][x] === 1) continue;

    // 가로 → 가로 or 대각선
    if (x - 1 >= 0) {
      dp[y][x][0] += dp[y][x - 1][0] + dp[y][x - 1][2];
    }

    // 세로 → 세로 or 대각선
    if (y - 1 >= 0) {
      dp[y][x][1] += dp[y - 1][x][1] + dp[y - 1][x][2];
    }

    // 대각선 → 세 방향 모두 가능
    if (x - 1 >= 0 && y - 1 >= 0 &&
        arr[y - 1][x] === 0 &&
        arr[y][x - 1] === 0) {
      dp[y][x][2] += dp[y - 1][x - 1][0] + dp[y - 1][x - 1][1] + dp[y - 1][x - 1][2];
    }
  }
}

console.log(dp[N-1][N-1].reduce((acc, curr) => acc+curr,0))