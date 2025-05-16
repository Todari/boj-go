const input = require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim().split("\n");
const N = Number(input[0])
const coords = input.slice(1).map(line => line.split(" ").map(Number))

// const getTriArea = (coords) => {
//   const rect = (Math.max(...coords.map(coord => coord[0])) - Math.min(...coords.map(coord => coord[0]))) * 
//   (Math.max(...coords.map(coord => coord[1])) - Math.min(...coords.map(coord => coord[1])))

  
//   const area = rect - 0.5*((Math.abs(coords[0][0] - coords[1][0])*Math.abs(coords[0][1] - coords[1][1]))
//   +(Math.abs(coords[1][0] - coords[2][0])*Math.abs(coords[1][1] - coords[2][1]))
//   +(Math.abs(coords[2][0] - coords[0][0])*Math.abs(coords[2][1] - coords[0][1])))
  
//   return area
// }

// let answer = 0;
// let last = 1
// for(let i = 2; i<N; i++) {
//   answer += getTriArea([coords[0], coords[last], coords[i]])
//   last = i
// }

// console.log(answer.toFixed(1))

let area = 0;
for (let i = 0; i < N; i++) {
  const [x1, y1] = coords[i];
  const [x2, y2] = coords[(i + 1) % N];
  area += (x1 * y2 - x2 * y1);
}

console.log((Math.abs(area) / 2).toFixed(1))