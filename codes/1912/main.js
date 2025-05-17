const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n");
const n = Number(input[0])
let dp = Array.from({length: n + 1}, () => -Infinity)
const arr = input[1].split(" ").map((number, idx) => {
  const num = Number(number)
  dp[idx + 1] = Math.max(num, dp[idx] + num)
  return num
})

console.log(Math.max(...dp))