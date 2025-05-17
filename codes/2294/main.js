const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n");
const [n, k] = input[0].split(" ").map(Number)
let dp = Array.from({length: k+1}, () => Infinity)
const coins = input.slice(1).map(coin => {
  dp[coin] = 1;
  return Number(coin)
})

for (const coin of coins) {
  for (let i=coin; i<=k; i++) {
    dp[i] = Math.min(dp[i], dp[i - coin] + 1)
  }
}

console.log(dp[k]===Infinity?-1:dp[k])