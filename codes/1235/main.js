const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString()

const n = Number(input.split("\n")[0]);
const numbers = input.split("\n").slice(1);
const length = numbers[0].length;

for (let i = 1; i <= length; i++) {
  let set = new Set();
  for (let j = 0; j < n; j++) {
    let newNumber = numbers[j].slice(length-i, length)
    set.add(newNumber)
  }
  
  if (set.size === n) {
    console.log(i)
    break;
  }
}