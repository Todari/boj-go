const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n");
const N = Number(input[0])
const lines = input.slice(1).map(line => line.split(" ").map(Number)).sort((a,b) => a[0] - b[0])

let dp = Array.from({length: N}, () => 1)

for (let i=1; i<N; i++) {
  for (let j=0; j<i; j++) {
    if (lines[j][1] < lines[i][1]) {
      dp[i] = Math.max(dp[i], dp[j] + 1);
    }
  }
}

console.log(N - Math.max(...dp))