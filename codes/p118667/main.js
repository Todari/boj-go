function solution(queue1, queue2) {
    // BigInt를 사용하여 값 저장
    let q1Sum = queue1.reduce((a, b) => a + BigInt(b), BigInt(0));
    let q2Sum = queue2.reduce((a, b) => a + BigInt(b), BigInt(0));
    let totalSum = q1Sum + q2Sum;

    // 전체 합이 홀수라면 불가능
    if (totalSum % BigInt(2) !== BigInt(0)) return -1;
    
    let target = totalSum / BigInt(2);
    let combined = [...queue1, ...queue2];  // 하나의 배열로 결합
    let left = 0, right = queue1.length;
    let sum = q1Sum;
    let limit = queue1.length * 3; // 최악의 경우를 대비한 제한

    for (let count = 0; count < limit; count++) {
        if (sum === target) return count;
        if (sum < target) {
            sum += BigInt(combined[right]); // 오른쪽 값을 더함
            right++;
        } else {
            sum -= BigInt(combined[left]); // 왼쪽 값을 뺌
            left++;
        }
        
        // 포인터가 배열 크기를 벗어나면 종료 (불가능한 경우)
        if (right >= combined.length) break;
    }

    return -1;
}

// function solution(queue1, queue2) {

//      // BigInt를 사용하여 값 저장
//     let q1Sum = queue1.reduce((a, b) => a + BigInt(b), BigInt(0));
//     let q2Sum = queue2.reduce((a, b) => a + BigInt(b), BigInt(0));
//     let totalSum = q1Sum + q2Sum;

//     // 전체 합이 홀수라면 불가능
//     if (totalSum % BigInt(2) !== BigInt(0)) return -1;
    
//     let target = totalSum / BigInt(2);
//     var loop = [...queue1, ...queue2]
//     let length = loop.length
//     let half = length / 2
    
//     let left = 0;
//     let right = 1;
//     let sum = BigInt(loop[0])
    
//     while (right < length && left < right) {
//         if (sum < target) {
//             sum += BigInt(loop[right])
//             right++;
//         } else if (sum > target) {
//             sum -= BigInt(loop[left])
//             left++;
//         } else {
//             break;
//         }
//     }
    
//     if (left >= half  && right > half) {
//         left -= half;
//         right -= half;
//     }
    
//     var answer = (right === length || left === right) ? -1 : right >= half ? left + right - half : left + right + half
    
//     return answer;
// }