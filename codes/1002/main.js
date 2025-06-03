const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n");
const T = Number(input[0])
for (let i = 0; i < T; i++) {
  const [x1, y1, r1, x2, y2, r2] = input[i + 1].split(" ").map(Number);
  const dx = x2 - x1;
  const dy = y2 - y1;
  const dSquared = dx * dx + dy * dy;

  const sumR = r1 + r2;
  const diffR = Math.abs(r1 - r2);
  const sumRSquared = sumR * sumR;
  const diffRSquared = diffR * diffR;

  if (dSquared === 0) {
    if (r1 === r2) {
      console.log(-1);
    } else {
      console.log(0);
    }
  } else if (dSquared === sumRSquared || dSquared === diffRSquared) {
    console.log(1); 
  } else if (diffRSquared < dSquared && dSquared < sumRSquared) {
    console.log(2); 
  } else {
    console.log(0); 
  }
}