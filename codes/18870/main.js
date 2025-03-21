const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n")
const N = Number(input[0])
let arr = input[1].split(" ").map(Number)
let set = new Set([...arr].sort((a,b) => a-b))
let map = new Map();

let counter = 0;
for (let value of set) {
  map.set(value, counter)
  counter++;
}

console.log(arr.map((v) => map.get(v)).join(" "))