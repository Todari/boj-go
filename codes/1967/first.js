const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n");
const n = Number(input[0])
const tree = Array.from({length: n + 1}, () => [])
const length = Array.from({length: n + 1}, () => [])
length[1] = [0, []]
for (const [parent, child, weight] of input.slice(1).map((str) => str.split(" ").map(Number))) {
  tree[parent] = [...tree[parent], {child, weight}]
}

let visited = [];

function dfs(start) {
  if (!tree[start] || visited.includes(start)) return
  visited.push(start)
  for(const {child, weight} of tree[start]) {
    length[child] = [length[start][0]+weight, [...length[start][1], child]]
    dfs(child)
  }
}

dfs(1)
let temp = [...length]
temp.sort((a, b) => b[0] - a[0])
let answer = temp[1][0] + temp[2][0]
let duplicated = 0;

for (let i = 0; i<temp[1][1].length; i++) {
  if (temp[1][1][i] === temp[2][1][i]) {
    duplicated = temp[1][1][i]
  } else {
    break;
  }
}

answer -= 2 * length[duplicated][0]

console.log(answer)