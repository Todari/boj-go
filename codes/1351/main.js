const [N, P, Q] = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split(" ").map(Number);
const map = new Map()
map.set(0, 1)
map.set(1, 2)

const getX = (number) => {
  if (map.has(number)) return map.get(number)

  const first = map.get(Math.floor(number/P)) ?? getX(Math.floor(number/P))
  const second = map.get(Math.floor(number/Q)) ?? getX(Math.floor(number/Q))
  map.set(number, first + second)
  return first + second
}

console.log(getX(N))