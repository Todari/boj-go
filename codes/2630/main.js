const fs = require("fs");

// 입력 받기
const input = fs.readFileSync("/dev/stdin").toString().trim();
const n = Number(input.split("\n")[0]);
const step = Math.log2(n);

let colors = [[]];
input.split("\n").slice(1).forEach((line) => {
    colors[0].push(...line.split(" ").map(Number));
});

// 정사각형이 모두 같은 색인지 확인하는 함수
const isCompleted = (array) => {
    return array.every((val) => val === array[0]);
};

// 정사각형을 4등분하는 함수
const cut = (array) => {
    let result = [[], [], [], []];
    let len = Math.sqrt(array.length);
    let half = len / 2;

    if (isCompleted(array)) return [array];

    for (let i = 0; i < len; i++) {
        for (let j = 0; j < len; j++) {
            let index = i * len + j;
            if (i < half && j < half) result[0].push(array[index]);
            else if (i < half && j >= half) result[1].push(array[index]);
            else if (i >= half && j < half) result[2].push(array[index]);
            else result[3].push(array[index]);
        }
    }

    return result;
};

// `step`만큼 반복하며 색종이를 자르기
for (let i = 0; i < step; i++) {
    let tmp = [];
    for (let j = 0; j < colors.length; j++) {
        tmp.push(...cut(colors[j]));
    }
    colors = tmp;
}

// 결과 저장
let result = new Map();
result.set("W", 0);
result.set("B", 0);

// 흰색(W)과 파란색(B) 개수 세기
colors.forEach((array) => {
    if (array.every((val) => val === 0)) {
        result.set("W", result.get("W") + 1);
    } else if (array.every((val) => val === 1)) {
        result.set("B", result.get("B") + 1);
    }
});

// 최종 출력
console.log(result.get("W"));
console.log(result.get("B"));