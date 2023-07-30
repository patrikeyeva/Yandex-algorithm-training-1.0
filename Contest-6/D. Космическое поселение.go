//https://contest.yandex.ru/contest/27844/problems/D/

package main

import "fmt"

func max(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

func main() {

	var n, a, b, w, h uint64
	var l, r uint64
	fmt.Scan(&n, &a, &b, &w, &h)
	l, r = 0, max(w, h)
	for l < r {
		d := (l + r + 1) / 2
		newA, newB := a+2*d, b+2*d
		if (w/newA)*(h/newB) >= n || (w/newB)*(h/newA) >= n {
			l = d
		} else {
			r = d - 1
		}

	}

	fmt.Print(l)

}
