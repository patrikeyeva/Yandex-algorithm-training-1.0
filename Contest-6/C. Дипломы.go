//https://contest.yandex.ru/contest/27844/problems/C/

package main

import "fmt"

func max(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

func main() {

	var w, h, n uint64
	var l, r uint64
	fmt.Scan(&w, &h, &n)
	l, r = 0, max(w, h)*(n/2+1)
	for l < r {
		m := (l + r) / 2
		if (m/w)*(m/h) >= n {
			r = m
		} else {
			l = m + 1
		}

	}

	fmt.Print(l)

}
