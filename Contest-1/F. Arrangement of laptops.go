// https://contest.yandex.ru/contest/27393/problems/F/?success=88641451#30404/2021_05_13/vBnD98HT49

package main

import (
 "fmt"
)

func main() {
 var a1, b1, a2, b2 int
 fmt.Scan(&a1, &b1, &a2, &b2)

 if max(a1, b1) > max(a2, b2) {
  a1, b1, a2, b2 = a2, b2, a1, b1
 }

 if b2 > a2 {
  b2, a2 = a2, b2
 }
 if b1 > a1 {
  b1, a1 = a1, b1
 }

 if (a2+b1)*max(b2, a1) < a2*(b2+b1) {
  fmt.Println(a2+b1, max(b2, a1))
 } else {
  fmt.Println(a2, b2+b1)
 }
}

func max(a, b int) int {
 if a > b {
  return a
 }
 return b
}
