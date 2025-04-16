const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n")
const [N, M] = input[0].split(" ").map(Number)

let board = input.slice(1, 1+N).map((line) => line.split(" ").map(Number))
let targets = input.slice(1+N).map((line) => line.split(" ").map(Number))

let len = board.length
let dp = Array.from({length: len + 1}, () => Array.from({length: len + 1}, () => 0))

dp[1][1] = board[0][0]
for (let i = 1; i<len+1; i++) {
  for (let j = 1; j<len+1; j++) {
    if (i===1 && j===1) continue
    let top = dp[i-1][j]
    let left = dp[i][j-1]
    dp[i][j] = top + left + board[i-1][j-1] - dp[i-1][j-1]
  }
}

for (const [x1,y1,x2,y2] of targets) {
  console.log(dp[x2][y2] - dp[x1 - 1][y2] - dp[x2][y1 - 1] + dp[x1 - 1][y1 - 1])
}