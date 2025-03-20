const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().split("\n")
let N = input[0]
let map = input.slice(1).map((line) => [...line.split("").map(Number)])
let visited = Array.from({length : N}, () => Array.from({length: N}).fill(false))
let directions = [[0, 1], [1, 0], [0, -1], [-1, 0]]
let groups = new Map();
let id = 1;

for (let i = 0 ; i< N ; i++) {
  for (let j = 0; j< N; j++) {
    if (visited[i][j]) continue;
    if (map[i][j] === 1) {
      let size = 0;
      let queue = [[i, j]];

      while(queue.length>0) {
        let [y,x] = queue.pop();
        if (visited[y][x]) continue
        visited[y][x] = true;
        size++;
        for (let dir of directions) {
          let ny = y + dir[0]
          let nx = x + dir[1]
          if (ny >= 0 && ny < N && nx >= 0 && nx < N) {
            if (map[ny][nx] === 1 && !visited[ny][nx]) {
              queue.push([ny, nx])
            }
          }
        }
      }
      groups.set(id, size)
      id++;
    }
  }

}

console.log(groups.size)
let sortedGroups = [...groups.values()].sort((a,b)=> a-b)
for (let i=0; i<groups.size; i++) {
  console.log(sortedGroups[i])
}