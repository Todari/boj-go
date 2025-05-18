const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n");
const N = Number(input[0])
const arr = input[1].split(" ").map(Number)

let dp = Array.from({length: N}, () => 1)
let prev = Array.from({length: N}, () => -1)

for (let i = 1; i<N; i++) {
  for (let j=0; j<=i; j++) {
    if (arr[i] > arr[j] && dp[i] < dp[j] + 1){
      dp[i] = dp[j]+1
      prev[i] = j;
    }
  }
}

const maxLen = Math.max(...dp)
const answer = [];
let idx = dp.indexOf(maxLen)
while(idx !== -1) {
  answer.unshift(arr[idx])
  idx = prev[idx];
}

console.log(maxLen)
console.log(answer.join(" "))