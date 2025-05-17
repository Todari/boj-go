const input = Number(require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim());
let dp = Array.from({length: 1001}, () => Array.from(10).fill(0))
dp[1] = [1, 1, 1, 1, 1, 1, 1, 1, 1, 1]

for (let i = 2; i<=1000; i++) {
  for (let j = 0; j<10; j++) {
      dp[i][j] = dp[i-1].slice(0, j+1).reduce((acc,curr) => (acc+curr) % 10007, 0)
  }
}

console.log(dp[input].reduce((acc,curr) => (acc+curr) % 10007, 0))

