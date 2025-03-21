const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n")
const T = Number(input[0])

for (let r = 0; r<T; r++) {
  let cmds = input[r * 3 + 1]
  let arrStr = input[r * 3 + 3]
  let arrSlice = arrStr.slice(1, arrStr.length - 1)
  let arr = arrSlice === "" ? [] : arrSlice.split(",").map(Number)
  let reversed = false;
  let err = false;

  for (let cmd of cmds) {
    if (err) break;
    if (cmd === "R") {
      reversed = !reversed
    }
    if (cmd === "D") {
      if (arr.length === 0) err = true;
      reversed? arr.pop() : arr.shift()
    }
  }
  console.log(err? "error" : reversed?`[${arr.reverse().join(",")}]` : `[${arr.join(",")}]`)
}