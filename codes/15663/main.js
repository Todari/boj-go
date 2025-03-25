const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n")
const [N, M] = input[0].split(" ").map(Number)
const arr = input[1].split(" ").map(Number).sort((a,b) => a - b)

function getPermutation(arr, select) {
  if (select === 1) return arr.map(v => [v])

  let result = [];
  arr.forEach((v, i, o) => {
    let fixed = v
    let rest = [...o.slice(0, i), ...o.slice(i+1)]
    let next = getPermutation(rest, select - 1)
    let attached = next.map((value) => [fixed, ...value])
    result.push(...attached)
  })

  return result
}

let set = new Set(getPermutation(arr, M).map(v => v.join(" ")))
console.log(Array.from(set).join("\n"))