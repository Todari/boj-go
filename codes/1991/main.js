const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n")
const n = Number(input[0])

let map = new Map();
input.slice(1).forEach((line) => {
  let [curr, left, right] = line.split(" ").map((elem) => elem!=="." ? elem : null)
  map.set(curr, {left, right})
})

let preOrderResult = ""
let inOrderResult = ""
let postOrderResult = ""

function preOrder(node) {
  if (!node) return;
  let {left, right} = map.get(node)
  preOrderResult += node
  preOrder(left)
  preOrder(right)
}

function inOrder(node) {
  if (!node) return;
  let {left, right} = map.get(node)
  inOrder(left)
  inOrderResult += node
  inOrder(right)
}

function postOrder(node) {
  if (!node) return;
  let {left, right} = map.get(node)
  postOrder(left)
  postOrder(right)
  postOrderResult += node
}

preOrder('A')
inOrder('A')
postOrder('A')

console.log(preOrderResult)
console.log(inOrderResult)
console.log(postOrderResult)