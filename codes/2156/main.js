const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n");
const n = Number(input[0])
const volumes = input.slice(1).map(Number)
let dp = Array.from({length: n + 1}, () => 0)
dp[1] = volumes[0]
dp[2] = dp[1] + volumes[1]

for (let i=3; i<=volumes.length; i++) {
  dp[i] = Math.max(dp[i-1], dp[i-2] + volumes[i-1], dp[i-3] + volumes[i-2] + volumes[i-1])
}

console.log(dp[n])