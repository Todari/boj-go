---
difficulty: s3
class: class3
link: https://www.acmicpc.net/problem/17626
---
### 기존  접근
dp 를 통해 n의 수를 만들 수 있는 dp[x] + dp[y]의 최솟값을 통해 계산

### 기존 코드
 
```go
package main  
  
import (  
 "bufio"  
 "fmt" "math" "os")  
  
func main() {  
 w := bufio.NewWriter(os.Stdout)  
 r := bufio.NewReader(os.Stdin)  
 defer w.Flush()  
  
 var n int  
 fmt.Fscanln(r, &n)  
  
 dp := make([]int, 50001)  
 for i := 1; i < 224; i++ {  
  dp[int(math.Pow(float64(i), float64(2)))] = 1  
 }  
 for i := 2; i < 50001; i++ {  
  if dp[i] == 0 {  
   dp[i] = i  
  }  
  min := dp[i]  
  
  for j := 1; j < i/2; j++ {  
   sum := dp[j] + dp[i-j]  
   if sum <= min {  
    min = sum  
    dp[i] = sum  
   }  
  }  
  
 }  
 fmt.Fprintln(w, dp[n])  
}
```

이중 for문으로 인하여 시간초과가 날 것이라고 예상함
예상대로 시간초과

### 개선  접근
dp[n]은 dp[n-x^2] 들 중 제일 작은 값에서 1을 더한 값과 같음

### 기존 코드
 
```go
package main  
  
import (  
 "bufio"  
 "fmt" "os")  
  
func main() {  
 w := bufio.NewWriter(os.Stdout)  
 r := bufio.NewReader(os.Stdin)  
 defer w.Flush()  
  
 var n int  
 fmt.Fscanln(r, &n)  
  
 dp := make([]int, n+1)  
 dp[1] = 1  
 for i := 2; i < n+1; i++ {  
  min := i  
  
  for j := 1; j*j <= i; j++ {  
   if min > dp[i-j*j] {  
    min = dp[i-j*j]  
   }  
  }  
  dp[i] = min + 1  
 }  
  
 //fmt.Fprintln(w, dp)  
 fmt.Fprintln(w, dp[n])  
}
```

모든 합의 경우를 찾는 경우 보다, 주어진 값에서 어떠한 수의 제곱수를 뺀 것에 1을 더한것과 값이 같기 떄문에 최소 경로를 효율적으로 찾을 수 있음
