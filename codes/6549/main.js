const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n");

for (const line of input) {
  const [n, ...arr] = line.split(" ").map(Number);
  if (n === 0) break;

  const heights = [...arr, 0];
  const stack = [];
  let maxArea = 0;

  for (let i = 0; i < heights.length; i++) {
    while (stack.length && heights[stack[stack.length - 1]] > heights[i]) {
      const height = heights[stack.pop()];
      const width = stack.length === 0 ? i : i - stack[stack.length - 1] - 1;
      maxArea = Math.max(maxArea, height * width);
    }
    stack.push(i);
  }

  console.log(maxArea);
}