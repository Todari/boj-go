const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n");
const [N, K] = input[0].split(" ").map(Number)

let dp = Array.from({length: 201}, () => Array.from({length: 201}, ()=>1))

for (let i=2; i<201; i++) {
  dp[i][0] = 1
  for (let j=1; j<201; j++) {
    dp[i][j] = (dp[i-1][j] + dp[i][j-1])%1000000000
  }
}

// console.log(dp)
console.log(dp[K][N])