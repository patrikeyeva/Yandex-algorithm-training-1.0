//https://contest.yandex.ru/contest/27472/problems/I/

package main

import (
	"fmt"
)

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	arr := make([]int, n*m)

	var p, q int
	for c := 0; c < k; c++ {
		fmt.Scan(&p, &q)
		p--
		q--
		index := p*m + q
		arr[index] = -1

		for i := max(0, p-1); i <= min(n-1, p+1); i++ {
			for j := max(0, q-1); j <= min(m-1, q+1); j++ {
				idx := i*m + j
				if arr[idx] != -1 {
					arr[idx]++
				}

			}
		}

	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			idx := i*m + j
			if arr[idx] == -1 {
				fmt.Printf("* ")
			} else {
				fmt.Printf("%d ", arr[idx])
			}
		}
		fmt.Printf("\n")
	}

}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
