const N = Number(require('fs').readFileSync(process.platform === "linux" ? "/dev/stdin" : "./input.txt").toString().trim())
const queue = Array.from({length:N}, (_,k) => [k])

// let answer = 0;

// while(queue.length) {
//   let selected = queue.shift()
//   if (selected.length === N) {
//     answer++;
//     continue
//   }
//   let next = Array.from({length:N}, () => true)
//   let len = selected.length
//   selected.forEach((v,k) => {
//     next[v] = false
//     if (0 <= v - (len-k)) next[v-(len-k)] = false
//     if (v + (len-k) < N) next[v+(len-k)] = false
//   })
//   next.forEach((v, k) => {if (v) queue.push([...selected, k])})
// }

// console.log(answer)

let answer = 0;
const col = Array(N).fill(false);
const diag1 = Array(2 * N).fill(false); // ↙ 방향 (row + col)
const diag2 = Array(2 * N).fill(false); // ↘ 방향 (row - col + N)

function dfs(row) {
  if (row === N) {
    answer++;
    return;
  }
  for (let c = 0; c < N; c++) {
    if (col[c] || diag1[row + c] || diag2[row - c + N]) continue;
    col[c] = diag1[row + c] = diag2[row - c + N] = true;
    dfs(row + 1);
    col[c] = diag1[row + c] = diag2[row - c + N] = false;
  }
}

dfs(0);
console.log(answer);