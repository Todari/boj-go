const input = require('fs').readFileSync(
  process.platform === "linux" ? "/dev/stdin" : "./input.txt"
).toString().trim().split("\n");

let strArr = input[1].split(''); // 입력 문자열을 문자 배열로 변환
let answer = 0;
const patterns = new Set([
  "skeep", "sXeep", "skXep", "skeXp", "skeeX",
  "sXXep", "sXeXp", "sXeeX", "skXeX", "skeXX",
  "sXXXp", "sXXeX", "sXeXX", "skXXX", "sXXXX"
]);

let stack = [];
for (let ch of strArr) {
  stack.push(ch);
  while (stack.length >= 5) {
    let lastFive = stack.slice(-5).join('');
    if (patterns.has(lastFive)) {
  
      for (let i = 0; i < 5; i++) stack.pop();
      stack.push("X");
      answer++;
      
    } else {
      break;
    }
  }
}

console.log(answer);