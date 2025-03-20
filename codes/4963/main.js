const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().split("\n");

let answers = [];
let directions = [[0, 1], [1, 1], [1, 0], [1, -1], [0, -1], [-1, -1] , [-1, 0], [-1, 1]];

for (let i = 0; i < input.length - 1; i++) {
    let [w, h] = input[i].split(" ").map(Number);
    if (w === 0 && h === 0) break;

    let map = [];
    for (let j = 0; j < h; j++) {
        map.push(input[i + j + 1].split(" ").map(Number));
    }

    let visited = Array.from({ length: h }, () => Array.from({ length: w }).fill(false));
    let count = 0;

    for (let y = 0; y < h; y++) {
        for (let x = 0; x < w; x++) {
            if (map[y][x] === 1 && !visited[y][x]) {
                let queue = [[y, x]];
                visited[y][x] = true;

                while (queue.length > 0) {
                    let [cy, cx] = queue.shift();

                    for (let [dy, dx] of directions) {
                        let ny = cy + dy;
                        let nx = cx + dx;

                        if (ny >= 0 && ny < h && nx >= 0 && nx < w && map[ny][nx] === 1 && !visited[ny][nx]) {
                            queue.push([ny, nx]);
                            visited[ny][nx] = true;
                        }
                    }
                }

                count++;
            }
        }
    }

    answers.push(count);
    i += h;
}

console.log(answers.join("\n"));