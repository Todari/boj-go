const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n")
const N = Number(input[0]);
const graph = input.slice(1).map(line => line.split(" ").map(Number))

for (let k=0; k<N; k++) {
    for (let i=0; i<N; i++) {
        for (let j=0; j<N; j++) {
            if (graph[i][k] === 1 && graph[k][j] === 1) {
                graph[i][j] = 1;
            }
        }
    }
}

console.log(graph.map(line => line.join(" ")).join("\n"))