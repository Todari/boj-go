const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n");
const N = Number(input[0])

let dp = Array.from({length: N + 1}, () => 0)
dp[0] = 1
dp[1] = 0
dp[2] = 3

for (let i = 3; i<=N; i++) {
  if (i%2 !==0) dp[i] = 0;
  else {
    dp[i] = dp[i-2] * 3 + dp[i-2] - dp[i-4]
  }
}

console.log(dp[N])