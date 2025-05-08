const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n");
let idx = 0;
const T = Number(input[idx++]);

for (let t = 0; t < T; t++) {
  const [N, M, W] = input[idx++].split(" ").map(Number);
  const edges = [];

  for (let i = 0; i < M; i++) {
    const [S, E, T] = input[idx++].split(" ").map(Number);
    edges.push([S, E, T]);
    edges.push([E, S, T]); // 양방향 도로
  }

  for (let i = 0; i < W; i++) {
    const [S, E, T] = input[idx++].split(" ").map(Number);
    edges.push([S, E, -T]); // 단방향 웜홀
  }

  console.log(bellmanFord(N, edges) ? "YES" : "NO");
}

function bellmanFord(N, edges) {
  const dist = Array(N + 1).fill(0); // 모든 노드에 대해 초기값 0으로 시작

  // N번 반복
  for (let i = 1; i <= N; i++) {
    for (const [u, v, w] of edges) {
      if (dist[v] > dist[u] + w) {
        dist[v] = dist[u] + w;

        // N번째 반복에서도 갱신된다면 음수 사이클 존재
        if (i === N) return true;
      }
    }
  }

  return false;
}