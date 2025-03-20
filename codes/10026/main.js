const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().split("\n");
const N = Number(input[0])
const map = input.slice(1).map((line) => line.split(""))
const directions = [[1, 0], [0, 1], [-1, 0], [0, -1]]

let noCBvisited = Array.from({length: N}, () => Array.from({length: N}).fill(false))
let CBvisited = Array.from({length: N}, () => Array.from({length: N}).fill(false))

let noColorBlind = 0;
let colorBlind = 0;

for (let y = 0; y<N; y++) {
  for (let x = 0; x<N; x++) {
    if (!noCBvisited[y][x]) {
      let queue  = [[y, x]]
      noCBvisited[y][x] = true;
      let color = map[y][x]
      let pointer = 0;

      while (pointer < queue.length) {
        let [cy, cx] = queue[pointer++]
        
        for (let [dy, dx] of directions) {
          let ny = cy + dy
          let nx = cx + dx

          if (0 <= ny && ny < N && 0 <= nx && nx < N && !noCBvisited[ny][nx] && map[ny][nx] === color) {
            noCBvisited[ny][nx] = true;
            queue.push([ny, nx])
          }
        }
      }
      noColorBlind++;
    }

    if (!CBvisited[y][x]) {
      let queue  = [[y, x]]
      CBvisited[y][x] = true;
      let color = map[y][x]
      let pointer = 0;

      while (pointer < queue.length) {
        let [cy, cx] = queue[pointer++]
        
        for (let [dy, dx] of directions) {
          let ny = cy + dy
          let nx = cx + dx

          if (0 <= ny && ny < N && 0 <= nx && nx < N && !CBvisited[ny][nx] && (color !== "B" ? map[ny][nx] === "R" || map[ny][nx] ==="G" : map[ny][nx] === "B")) {
            CBvisited[ny][nx] = true;
            queue.push([ny, nx])
          }
        }
      }
      colorBlind++;
    }
  }  
}
console.log(noColorBlind, colorBlind)