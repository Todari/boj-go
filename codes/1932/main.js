const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n")
const n = Number(input[0])
const map = input.slice(1).map((v) => v.split(" ").map(Number))

let dp = Array.from({length: n}, (v,k)=>Array.from({length:k+1}))
dp[0][0] = map[0][0]
for (let i = 1; i<n; i++) {
  for (let j = 0; j<i+1; j++) {
    dp[i][j] = Math.max(dp[i-1][j] ?? 0, dp[i-1][j-1] ?? 0) + map[i][j]
  }
}

console.log(Math.max(...dp[n-1]))