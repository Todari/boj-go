const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n")
let s = input[0].split("")
let commands = input.slice(2)
let stack = [];

for (let cmd of commands) {
  if (cmd === "L" && s.length > 0) {
    stack.push(s.pop())
  }
  if (cmd === "D" && stack.length > 0) {
    s.push(stack.pop())
  }
  if (cmd === "B" && s.length > 0) {
    s.pop()
  }
  if (cmd[0] === "P") {
    s.push(cmd.split(" ")[1])
  }
}

s.push(...stack.reverse())

console.log(s.join(""))