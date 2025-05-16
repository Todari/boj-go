const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n");
const [N, R, Q] = input[0].split(" ").map(Number)
const graph = Array.from({length: N + 1}, () => [])
const edges = input.slice(1, N).map(line=>{
  const arr = line.split(" ").map(Number)
  graph[arr[0]].push(arr[1])
  graph[arr[1]].push(arr[0])
  return arr
})

let subTreeSize = Array.from({length: N + 1}, () => 0)

const dfs = (node, parent) => {
  subTreeSize[node] = 1
  for (const child of graph[node]) {
    if (child === parent) continue;
    dfs(child, node)
    subTreeSize[node] += subTreeSize[child]
  }
}

dfs(R, -1)

input.slice(N).forEach(line => {
  console.log(subTreeSize[Number(line)])
})