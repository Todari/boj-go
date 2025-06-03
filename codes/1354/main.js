const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split(" ").map(Number);
const [N, P, Q, X, Y] = input
const map = new Map();
map.set(0, 1)

const getA = (number) => {
  if (number <= 0) return 1

  let firstKey = Math.floor(number/P) - X
  if (firstKey <= 0) firstKey = 0
  let secondKey = Math.floor(number/Q) - Y
  if (secondKey <= 0) secondKey = 0

  const first = map.get(firstKey) ?? getA(firstKey)
  const second = map.get(secondKey) ?? getA(secondKey)

  map.set(number, first+second)
  return first+second
}

console.log(getA(N))