//https://contest.yandex.ru/contest/27844/problems/G/

package main

import (
	"fmt"
)

func min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

func main() {

	var n, m, t, l, r uint64
	fmt.Scan(&n, &m, &t)

	l, r = 0, min(n, m)/2
	for l < r {
		width := (l + r + 1) / 2
		if width*m*2+width*n*2-4*width*width <= t {
			l = width
		} else {
			r = width - 1
		}
	}

	fmt.Print(l)

}
