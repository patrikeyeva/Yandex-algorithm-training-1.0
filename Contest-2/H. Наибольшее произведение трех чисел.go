//https://contest.yandex.ru/contest/27472/problems/H/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var max_p, max_n, min_n []int

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())

		switch num >= 0 {
		case true:
			if len(max_p) < 3 {
				max_p = append(max_p, num)
				numsSort(&max_p)
			} else if abs(num) > abs(max_p[2]) {
				max_p[2] = num
				numsSort(&max_p)
			}

		case false:
			if len(max_n) < 2 {
				max_n = append(max_n, num)
				numsSort(&max_n)
			} else if abs(num) > abs(max_n[1]) {
				max_n[1] = num
				numsSort(&max_n)
			}
			if len(min_n) < 3 {
				min_n = append(min_n, num)
				sort.Ints(min_n)
			} else if abs(num) < abs(min_n[0]) {
				min_n[0] = num
				sort.Ints(min_n)
			}
		}
	}

	switch len(max_p) {
	case 0:
		fmt.Printf("%d %d %d", min_n[0], min_n[1], min_n[2])
	case 1:
		fmt.Printf("%d %d %d", max_p[0], max_n[0], max_n[1])
	case 2:
		if len(max_n) < 2 {
			fmt.Printf("%d %d %d", max_p[0], max_p[1], max_n[0])
		} else {
			fmt.Printf("%d %d %d", max_p[0], max_n[0], max_n[1])
		}
	case 3:
		if len(max_n) != 2 || max_p[0]*max_p[1]*max_p[2] > max_p[0]*max_n[0]*max_n[1] {
			fmt.Printf("%d %d %d", max_p[0], max_p[1], max_p[2])
		} else {
			fmt.Printf("%d %d %d", max_p[0], max_n[0], max_n[1])
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func numsSort(arr *[]int) {
	n := len(*arr)
	for i := n - 1; i > 0; i-- {
		if abs((*arr)[i-1]) >= abs((*arr)[i]) {
			break
		}
		(*arr)[i], (*arr)[i-1] = (*arr)[i-1], (*arr)[i]

	}
}
