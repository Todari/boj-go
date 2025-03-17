function solution(k, tangerine) {
    let box = new Map();
    for (const t of tangerine) {
        if (!box.has(t)) {
            box.set(t, 1)
        } else {
            box.set(t, box.get(t) + 1)
        }
    }
    
    let amounts = [...box.values()].sort((a, b) => b-a)
    
    let remain = k
    let answer = 0;
    for (let i = 0 ; i < amounts.length; i++) {
        
        remain -= amounts[i]
        if (remain <= 0) {
            answer = i + 1
            break;
        };
    }
    
    return answer
}