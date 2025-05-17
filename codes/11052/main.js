const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n");
const N = Number(input[0])
const prices = [0, ...input[1].split(" ").map(Number)]

let dp = [...prices]

for (let i=0; i<=N; i++) {
  for (let j=0; j<i; j++) {
    dp[i] = Math.max(dp[i], dp[i - j] + prices[j])
  }
}

console.log(dp[N])