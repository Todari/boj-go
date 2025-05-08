const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n")
const TC = Number(input[0])
let inputCounter = 1

for (let t = 0; t< TC; t++) {
  const [N, M, W] = input[inputCounter].split(" ").map(Number)
  const edges = [];
  let returned = false;
  inputCounter++
  for (let s = 0; s<M+W; s++) {
    const [S, E, T] = input[inputCounter].split(" ").map(Number)
    s<M ? edges.push([S,E,T]) : edges.push([S,E,-T])
    inputCounter++
  }
  
  for (let i = 1; i<=N; i++) {
    if (bellmanFord(N, edges, i)) {
      console.log("YES")
      returned = true;
      break;
    }
  }
  if (!returned) console.log("NO")
}

function bellmanFord(N, edges, start) {
  const dist = Array(N + 1).fill(Infinity);
  dist[start] = 0;

  for (let i = 1; i <= N - 1; i++) {
    for (const [u, v, w] of edges) {
      if (dist[u] !== Infinity && dist[v] > dist[u] + w) {
        dist[v] = dist[u] + w;
      }
    }
  }

  for (const [u, v, w] of edges) {
    if (dist[u] !== Infinity && dist[v] > dist[u] + w) {
      return true;
    }
  }

  return false
}