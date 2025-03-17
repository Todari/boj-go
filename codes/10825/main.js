const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString()

// const n = Number(input.split("\n")[0]);
const students = input.split("\n").slice(1).map((string) => string.split(" "));

const compare = (a, b) => {
  for (let i=0; i<a.length; i++) {
    if (b[i] === undefined) {
      return 1
    } else if (a[i] !== b[i]) {
      return a[i].charCodeAt(0) - b[i].charCodeAt(0)
    }
  }
  if ((a.length === b.length) && (a[a.length-1] === b[b.length-1])) {
    return 0
  } else {
    return -1
  }
}

students.sort((a,b) => {
  if (a[1] !== b[1]) {
    return Number(b[1]) - Number(a[1])
  } else if (a[2] !== b[2]) {
    return Number(a[2]) - Number(b[2])
  } else if (a[3] !== b[3]) {
    return Number(b[3]) - Number(a[3])
  } else {
    return compare(a[0], b[0])
  }
})

console.log(students.map((student) => student[0]).join("\n"))
