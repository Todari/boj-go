const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n")
const T = Number(input[0])
for (let i = 0; i< T; i++) {
  const n = Number(input[i*3 + 1])
  const sticker = input.slice(i*3 + 2, i*3 + 4).map((line) => line.split(" ").map(Number))

  let dp = Array.from({length: 2}, () => Array.from({length:n}))
  dp[0][0] = sticker[0][0]
  dp[1][0] = sticker[1][0]
  dp[0][1] = dp[1][0] + sticker[0][1]
  dp[1][1] = dp[0][0] + sticker[1][1]
  for (let i = 2; i<n; i++) {
    dp[0][i] = Math.max(dp[1][i-1] + sticker[0][i], dp[1][i-2] + sticker[0][i])
    dp[1][i] = Math.max(dp[0][i-1] + sticker[1][i], dp[0][i-2] + sticker[1][i])
  }
  
  console.log(Math.max(dp[0][n-1],dp[1][n-1]))
}