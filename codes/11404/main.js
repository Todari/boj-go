const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n")
const n = Number(input[0])
const m = Number(input[1])
let map = Array.from({length: n}, () => Array.from({length: n}, () => Infinity))
for (const [a, b, c] of input.slice(2).map((v) => v.split(" ").map(Number))) {
  map[a - 1][b - 1] = Math.min(map[a - 1][b - 1], c)
}

for (let k = 0; k < n; k++) {
  for (let i = 0; i < n; i++) {
    for (let j = 0; j < n; j++) {
      if (map[i][j] > map[i][k] + map[k][j]) {
        map[i][j] = map[i][k] + map[k][j];
      }
    }
  }
}

for (let i = 0; i < n; i++) {
  console.log(map[i].map((v, k) => v === Infinity || i===k ? 0 : v).join(" "));
}