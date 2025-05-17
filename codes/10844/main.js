const input = Number(require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim());
let dp = Array.from({length: 101} ,() => Array.from({length: 10}, () => 0))
dp[1] = [0, 1, 1, 1, 1, 1, 1, 1, 1, 1]
for (let i = 2; i<dp.length; i++) {
  for (let j = 0; j<10; j++) {
    if (j===0) {
      dp[i][j] = dp[i-1][j+1] %1000000000
    } else if (j===9) {
      dp[i][j] = dp[i-1][j-1] %1000000000
    } else {
      dp[i][j] = (dp[i-1][j-1] + dp[i-1][j+1])%1000000000
    }
  }
}

console.log(dp[input].reduce((acc,curr) => (acc + curr)%1000000000, 0))