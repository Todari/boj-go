const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim();
const N = Number(input)

let board = Array.from({ length: N }, () => Array(N).fill(' '));

function fillStar(x, y, size) {
  if (size === 1) {
    board[x][y] = '*';
    return;
  }

  const next = size / 3;

  for (let i = 0; i < 3; i++) {
    for (let j = 0; j < 3; j++) {
      if (i === 1 && j === 1) continue;
      fillStar(x + i * next, y + j * next, next);
    }
  }
}

fillStar(0, 0, N);

console.log(board.map((row) => row.join("")).join("\n"));