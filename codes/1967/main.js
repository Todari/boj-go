const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n");
const n = Number(input[0])
const tree = Array.from({length: n + 1}, () => [])

for (const [parent, child, weight] of input.slice(1).map((str) => str.split(" ").map(Number))) {
  tree[parent].push([child, weight])
  tree[child].push([parent, weight])
}

let max = 0;
let farthest = 0;

// console.log(tree)

function dfs(node, sum, visited) {
  visited[node] = true;
  if (sum > max) {
    max = sum;
    farthest = node;
  }

  for (let [next, dist] of tree[node]) {
    if (!visited[next]) {
      dfs(next, sum + dist, visited);
    }
  }
}

// 1. 임의의 노드에서 가장 먼 노드 찾기
dfs(1, 0, Array(n + 1).fill(false));

// 2. 그 노드에서 다시 가장 먼 노드까지 거리 계산
max = 0;
dfs(farthest, 0, Array(n + 1).fill(false));

console.log(max);