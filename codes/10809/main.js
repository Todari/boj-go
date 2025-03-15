const fs = require("fs");

// 입력 받기
const input = fs.readFileSync(0).toString().trim();

// a부터 z까지 각 알파벳의 위치 찾기
let result = [];
for (let i = 97; i <= 122; i++) {
  const char = String.fromCharCode(i);
  result.push(input.indexOf(char));
}

// 결과 출력
console.log(result.join(' '));