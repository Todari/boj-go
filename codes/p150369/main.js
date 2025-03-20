function solution(cap, n, deliveries, pickups) {
    let answer = 0;
    let lastDeliver = n - 1;
    let lastPickup = n - 1;

    while (lastDeliver >= 0 || lastPickup >= 0) {
        // 🚀 1. 가장 먼 배달 & 수거 위치 찾기
        while (lastDeliver >= 0 && deliveries[lastDeliver] === 0) lastDeliver--;
        while (lastPickup >= 0 && pickups[lastPickup] === 0) lastPickup--;
        
        let maxDist = Math.max(lastDeliver, lastPickup) + 1; // 🚚 이동 거리 계산 (왕복)

        let dCap = cap, pCap = cap;

        // 🚀 2. 먼 곳부터 배달 처리
        for (let i = lastDeliver; i >= 0 && dCap > 0; i--) {
            if (deliveries[i] > dCap) {
                deliveries[i] -= dCap;
                dCap = 0;
            } else {
                dCap -= deliveries[i];
                deliveries[i] = 0;
                lastDeliver--;
            }
        }

        // 🚀 3. 먼 곳부터 수거 처리
        for (let i = lastPickup; i >= 0 && pCap > 0; i--) {
            if (pickups[i] > pCap) {
                pickups[i] -= pCap;
                pCap = 0;
            } else {
                pCap -= pickups[i];
                pickups[i] = 0;
                lastPickup--;
            }
        }

        answer += maxDist * 2; // 왕복 거리 추가
    }

    return answer;
}

// function solution(cap, n, deliveries, pickups) {
//     var answer = 0;
//     var house = [];
    
//     for (let i = 0; i<n; i++) {
//         let tmp = [];
//         tmp.push(deliveries[i])
//         tmp.push(pickups[i])
//         house.push(tmp)
//     }
    
//     let sum = house.reduce((prev, curr) => prev[0] + prev[1] + curr[0] + curr[1] , 0)
//     let pointer = n - 1;
    
//     while (sum !== 0) {
//       if (pointer === 0) break;
//         if (house[pointer][0] + house[pointer][1] === 0) {
//             pointer --;
//             continue;
//         }
        
//         let deliveriesPointer = pointer;
//         let pickupsPointer = pointer;
//         let remainDeliveries = cap;
//         let remainPickups = cap;
        
//         while (remainDeliveries > 0) {
//           if (house.reduce((prev, curr) => prev + curr[0] , 0) === 0) break;
//             if (house[deliveriesPointer][0] === 0 && deliveriesPointer !== 0) {
//               deliveriesPointer-- ;
//             }
//             if (remainDeliveries < house[deliveriesPointer][0]) {
//               house[deliveriesPointer][0] -= remainDeliveries
//               remainDeliveries = 0
//             } else {
//               remainDeliveries -= house[deliveriesPointer][0]
//               house[deliveriesPointer][0] = 0
//               if (deliveriesPointer !== 0) deliveriesPointer--;
//             }
//         }
        
//         while (remainPickups > 0) {
//           if (house.reduce((prev, curr) => prev + curr[1] , 0) === 0) break;
//             if (house[pickupsPointer][1] === 0 && pickupsPointer !== 0) {
//               pickupsPointer-- ;
//             }
//             if (remainPickups < house[pickupsPointer][1]) {
//               house[pickupsPointer][1] -= remainPickups
//               remainPickups = 0
//             } else {
//               remainPickups -= house[pickupsPointer][1]
//               house[pickupsPointer][1] = 0
//               if (pickupsPointer !== 0) pickupsPointer--;
//             }
//         }
        
//         answer += 2 * pointer + 2
//     }

//     return answer;
// }

// solution(4, 5, [1,0,3,1,2], [0,3,0,4,0])
// solution(2, 7, [1, 0, 2, 0, 1, 0, 2],	[0, 2, 0, 1, 0, 2, 0])