//https://contest.yandex.ru/contest/27472/problems/J/

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	var a, b, num float64
	var word string
	fmt.Scan(&a)

	var min_res, max_res float64 = 30.0, 4000.0

	for i := 0; i < n-1; i++ {
		fmt.Scan(&num, &word)
		switch word {
		case "closer":
			b = num
		case "further":
			a, b = num, a

		}
		if b < a {
			max_res = math.Min(max_res, (a+b)/2)

		} else if b > a {
			min_res = math.Max(min_res, (a+b)/2)

		}
		a = num
	}

	fmt.Printf("%.6f %.6f", min_res, max_res)

}
