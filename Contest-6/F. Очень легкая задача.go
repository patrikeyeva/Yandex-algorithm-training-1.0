//https://contest.yandex.ru/contest/27844/problems/F/

package main

import (
	"fmt"
)

func main() {

	var n, x, y int
	fmt.Scan(&n, &x, &y)
	if y < x {
		x, y = y, x
	}

	l, r := 0, (x+y)*n
	for l < r {
		time := (l + r) / 2
		if (time/x)+(time/y) >= n-1 {
			r = time
		} else {
			l = time + 1
		}
	}

	fmt.Print(l + x)

}
