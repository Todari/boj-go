const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split(" ")
const [A, B] = input.map(Number)

const queue = [[A, 1]];

while (queue.length) {
  const [now, cnt] = queue.shift();

  if (now === B) {
    console.log(cnt);
    return;
  }

  if (now * 2 <= B) {
    queue.push([now * 2, cnt + 1]);
  }
  if (now * 10 + 1 <= B) {
    queue.push([now * 10 + 1, cnt + 1]);
  }
}

console.log(-1);