const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n");
const T = Number(input[0])

const answer = [];
for(let i=0; i<T ; i++) {
  const arr = input[i+1]
  const stack = [];
  for (const char of arr) {
    if (char==="(") {
      stack.push(char)
    } else if (stack.length === 0) {
      answer.push("NO")
      break;
    } else {
      stack.pop();
    }
  }

  if (answer.length === i+1) continue;
  if (stack.length === 0) answer.push("YES")
  else answer.push("NO")
}

console.log(answer.join("\n"))