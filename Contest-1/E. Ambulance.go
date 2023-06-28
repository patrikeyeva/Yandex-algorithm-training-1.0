//https://contest.yandex.ru/contest/27393/problems/E/

package main

import (
 "fmt"
 "math"
)

func main() {
 var K1, M, K2, P2, N2 int
 fmt.Scan(&K1, &M, &K2, &P2, &N2)
 P1 := -1
 N1 := -1

 // противоречия
 if M < N2 || M*(P2-1)+N2 > K2 {
  P1 = -1
  N1 = -1
 } else if N2 == 1 && P2 == 1 {
  if K1 <= K2 {
   P1 = 1
   N1 = 1
  } else if K1 <= M {
   P1 = 1
   N1 = 0
   if M == 1 {
    N1 = 1
   }
  } else {
   P1 = 0
   N1 = 0
   if M == 1 {
    N1 = 1
   }
  }
 } else {
  x1 := int(math.Ceil(float64(K2) / float64(N2+M*(P2-1))))
  x2 := float64(K2) / float64(N2-1+M*(P2-1))
  if float64(x1) < x2 {
   if x2 == float64(int(x2)) {
    x2 -= 1
   } else {
    x2 = math.Floor(x2)
   }

   P1 = int(math.Ceil(float64(K1) / (float64(x1) * float64(M))))
   N1 = int(math.Ceil((float64(K1) - float64(x1)*float64(M)*(float64(P1)-1)) / float64(x1)))
   checkP := true
   checkN := true
   for x := int(x1) + 1; x <= int(x2); x++ {
    P := int(math.Ceil(float64(K1) / (float64(x) * float64(M))))
    N := int(math.Ceil((float64(K1) - float64(x)*float64(M)*(float64(P)-1)) / float64(x)))
    if checkP && P != P1 {
     P1 = 0
     checkP = false
    }
    if checkP {
     P1 = P
    }

    if checkN && N != N1 {
     N1 = 0
     if M == 1 {
      N1 = 1
     }
     checkN = false
    }
    if checkN {
     N1 = N
    }
   }
  }
 }

 fmt.Println(P1, N1)
}
