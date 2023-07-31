//https://contest.yandex.ru/contest/27844/problems/E/

package main

import (
	"fmt"
)

func main() {

	var a, b, c uint64
	var l, r uint64
	fmt.Scan(&a, &b, &c)
	l, r = 0, 3*a+b+c+2
	for l < r {
		m := (l + r) / 2
		if 2*(2*a+3*b+4*c+5*m) >= 7*(a+b+c+m) {
			r = m
		} else {
			l = m + 1
		}

	}
	fmt.Print(l)

}
