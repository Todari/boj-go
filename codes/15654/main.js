const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n")
const [N, M] = input[0].split(" ").map(Number)
const arr = input[1].split(" ").map(Number).sort((a,b) => a-b)

function getPermutation(arr, select) {
  if (select === 1) return arr.map((v) => [v])

  let result = [];
  arr.forEach((value, index, origin) => {
    let fixed = value;
    let rest = [...origin.slice(0, index), ...origin.slice(index + 1)]
    let comb = getPermutation(rest, select - 1)
    let attached = comb.map((v) => [fixed, ...v])

    result.push(...attached)
  })
  return result
}

console.log(getPermutation(arr, M).map((v) => v.join(" ")).join("\n"))
