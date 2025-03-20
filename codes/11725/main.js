const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().split("\n")
const N = Number(input[0])

let graph = new Map();
let visited = Array.from({length: N + 1}).fill(false)
let answer = new Map();

for (let node of input.slice(1)) {
  let [a, b] = node.split(" ").map(Number)
  if (!graph.has(a)) graph.set(a, [])
  if (!graph.has(b)) graph.set(b, [])
  let aArr = graph.get(a)
  aArr.push(b)
  let bArr = graph.get(b)
  bArr.push(a)
  graph.set(a, aArr)
  graph.set(b, bArr)
}

function dfs(start) {
  visited[start] = true
  for (const next of graph.get(start)) {
    if (visited[next]) continue;
    answer.set(next, start);
    dfs(next)
  }
}

dfs(1)

for (let i = 2; i < N + 1 ; i++) {
  console.log(answer.get(i))
}