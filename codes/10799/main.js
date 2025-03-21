const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim()

let answer = 0;

let stack = [];
let sticks = input.replaceAll("()", "X").split("")

for (let char of sticks) {
  if (char === "X") {
    answer += stack.length
  }
  if (char === "(") {
    stack.push(char)
  }
  if (char === ")") {
    stack.pop()
    answer++;
  }
}

console.log(answer)