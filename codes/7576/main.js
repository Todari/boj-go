const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().split("\n");
const [M, N] = input[0].split(" ").map(Number)
const directions = [[1, 0], [0, 1], [-1, 0], [0, -1]]

let tomatoes = Array.from({length: N}, () => Array.from({length: M}).fill(0))
let days = Array.from({length: N}, () => Array.from({length: M}).fill(Infinity))
let queue = []
let front = 0;

input.slice(1).forEach((line, i) => line.split(" ").forEach((v, j) => {
	if (v==="1") {
		queue.push([i,j,0])
		days[i][j] = 0;
		tomatoes[i][j] = 1
	} else if (v==="-1") {
		days[i][j] = 0;
		tomatoes[i][j] = -1
	}
}))

while (front < queue.length) {
	let [cy, cx, day] = queue[front++];

	for (let [dy, dx] of directions) {
		let ny = cy + dy
		let nx = cx + dx

		if (0 <= ny && ny < N &&  0 <= nx && nx < M && tomatoes[ny][nx] === 0) {
			if (days[ny][nx] === Infinity) {
				days[ny][nx] = day + 1
				queue.push([ny, nx, day + 1])
			}
		}
	}
}

const max = Math.max(...days.flat())

console.log(max === Infinity? -1 : max)