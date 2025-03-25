const [N, M] = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split(" ").map(Number)

let result = [];
let sequence = [];

function backtrack(start, depth) {
  if (depth === M) {
    result.push(sequence.join(" "));
    return;
  }

  for (let i = start; i <= N; i++) {
    sequence.push(i);
    backtrack(i, depth + 1);
    sequence.pop();
  }
}

backtrack(1, 0);
console.log(result.join("\n"));

// const [N, M] = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split(" ").map(Number)

// function getCombination(arr, select) {
//   if (select === 1) return arr.map((v) => [v])

//   let result = [];
//   arr.forEach((value, index, origin) => {
//     let fixed = value;
//     let rest = origin.slice(index)
//     let comb = getCombination(rest, select - 1)
//     let attached = comb.map((v) => [fixed, ...v])

//     result.push(...attached)
//   })
//   return result
// }

// let arr = Array.from({length: N}, (_, k) => k + 1)
// console.log(getCombination(arr, M).map(v => v.join(" ")).join("\n"))
