const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().split("\n")

const numbers = input.map((v, i) => [i, v]).filter((value) => !isNaN(value[1]))

let result = ""
if ((Number(numbers[numbers.length - 1][1]) + 3 - Number(numbers[numbers.length - 1][0])) % 3 === 0) {
  result += "Fizz"
}
if ((Number(numbers[numbers.length - 1][1]) + 3 - Number(numbers[numbers.length - 1][0])) % 5 === 0) {
  result += "Buzz"
}
if (result.length === 0) {
  result = `${Number(numbers[numbers.length - 1][1]) + 3 - Number(numbers[numbers.length - 1][0])}`
}
console.log(result)