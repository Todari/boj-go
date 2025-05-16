const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n");
const N = Number(input[0])
const attributes = input[1].split(" ").map(Number)

let left = 0;
let right = attributes.length - 1
let min = Infinity
let answer = [attributes[left], attributes[right]];

while (left < right) {
  const sum = attributes[left] + attributes[right];

  if (Math.abs(sum) < Math.abs(min)) {
    min = sum;
    answer = [attributes[left], attributes[right]];
  }

  if (sum < 0) left++;
  else right--;
}

console.log(answer.join(' '));