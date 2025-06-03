const N = Number(require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim());
let dp = Array.from({length: N + 1}, () => [0, 0, 0])
dp[1] = [1, 1, 1]

for (let i = 2; i<N + 1; i++) {
  dp[i][0] = (dp[i-1][0] + dp[i-1][1] + dp[i-1][2]) % 9901
  dp[i][1] = (dp[i-1][0] + dp[i-1][2]) % 9901
  dp[i][2] = (dp[i-1][0] + dp[i-1][1]) % 9901
}

console.log((dp[N][0] + dp[N][1] + dp[N][2])% 9901)

