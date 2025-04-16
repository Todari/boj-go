const [A,B,C] = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split(" ").map(BigInt)

function power(a, b) {
  if (b === 0n) return 1n;  
  if (b === 1n) return a % C;

  let half = power(a, b / 2n);
  let result = (half * half) % C;

  if (b % 2n === 1n) {
    result = (result * a) % C;
  }

  return result;
}

// BigInt의 n을 제거하기 위해 toString 사용
console.log(power(A, B).toString());