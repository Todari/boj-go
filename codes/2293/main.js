const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n");
const [n,k] = input[0].split(" ").map(Number)
const coins = input.slice(1).map(Number)
let dp = Array.from({length: k + 1}, () => 0)
dp[0] = 1

for (const coin of coins) {
  for (let i = coin; i<=k; i++) {
    dp[i] += dp[i - coin]
  }
}

console.log(dp[k])