const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().split("\n")
const N = Number(input[0]);
const graph = Array.from({ length: N }, () => []);

let answer = Array.from({ length: N }, () => Array(N).fill(0));

input.slice(1).forEach((row, rIdx) => {
    row.split(" ").forEach((x, cIdx) => {
        if (x === "1") {
            graph[rIdx].push(cIdx);
        }
    });
});

function dfs(start, node, visited) {
    for (let next of graph[node]) {
        if (!visited.has(next)) {
            visited.add(next);
            answer[start][next] = 1;
            dfs(start, next, visited);
        }
    }
}

// 각 정점에서 DFS 실행
for (let i = 0; i < N; i++) {
    let visited = new Set();
    dfs(i, i, visited);
}

// 결과 출력
console.log(answer.map(row => row.join(" ")).join("\n"));


// 플로이드-워셜 알고리즘
// for (let k = 0; k < N; k++) {
//     for (let i = 0; i < N; i++) {
//         for (let j = 0; j < N; j++) {
//             if (graph[i][k] && graph[k][j]) {
//                 graph[i][j] = 1;
//             }
//         }
//     }
// }

// // 결과 출력
// console.log(graph.map(row => row.join(" ")).join("\n"));