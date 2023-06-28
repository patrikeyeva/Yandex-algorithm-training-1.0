// https://contest.yandex.ru/contest/27393/problems/H/

package main

import "fmt"

func main() {
    var a, b, n, m int
    fmt.Scan(&a, &b, &n, &m)

    f1 := (n - 1) * a + n
    f2 := (n + 1) * a + n

    s1 := (m - 1) * b + m
    s2 := (m + 1) * b + m

    if s1 > f2 || f1 > s2 {
        fmt.Println(-1)
    } else {
        fmt.Println(max(f1, s1), min(f2, s2))
    }
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
