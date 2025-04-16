const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n")
let [short, long] = input.sort((a,b) => a.length - b.length)

let dp = Array.from({length: short.length + 1}, () => Array.from({length: long.length + 1}, () => 0))

for (let i = 1; i<short.length + 1; i++) {
  for (let j = 1; j<long.length + 1; j++) {
    if (short[i-1] === long[j-1]) {
      dp[i][j] = dp[i-1][j-1] + 1;
    } else {
      dp[i][j] = Math.max(dp[i-1][j], dp[i][j-1])
    }
  }
}

console.log(dp[short.length][long.length])