const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString()
const n = Number(input.split("\n")[0]);
var ropes = input.split("\n").map(Number).slice(1);

// 내림
// [60, 20, 15]
var sortedRopes = ropes.sort((a, b) => b - a);
// 오름
// [15, 20, 60]
// var sortedRopes = ropes.sort((a, b) => a - b);

var answer = 0;

for (let i = 0; i < n; i++) {
  //내림차순
    answer = Math.max(answer, sortedRopes[i] * (i + 1))
  //오름차순
    // answer = Math.max(answer, sortedRopes[i] * (n - i))
}

console.log(answer)