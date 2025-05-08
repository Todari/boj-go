const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim();
const N = Number(input)

let result = [];
let count = 0;

function hanoi(n, from, via, to) {
  if (n === 1) {
    result.push(`${from} ${to}`);
    count++;
    return;
  }

  hanoi(n - 1, from, to, via);
  result.push(`${from} ${to}`);
  count++;
  hanoi(n - 1, via, from, to);
}

hanoi(N, 1, 2, 3);
console.log(count);
console.log(result.join("\n"));

// move(3, 1, 3)

// move(2, 1, 2)
// 1 3
// move(2, 2, 3)

// move(1, 1, 3)
// 1 2
// move(1, 3, 1)
// 1 3
// move(1, 2, 1)
// 2 3
// move(1, 1, 3)

// 1 3


// 1 2

// 1 3
// 2 3


// 1 3

// 1 2
// 3 2

// 1 3
// 2 1
// 2 3
// 1 3


// 1 2

// 1 3
// 2 3

// 1 2
// 3 1
// 3 2
// 1 2

// 1 3
// 2 3
// 2 1
// 3 1
// 2 3
// 1 2
// 1 3
// 2 3