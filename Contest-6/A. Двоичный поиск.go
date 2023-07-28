//https://contest.yandex.ru/contest/27844/problems/A/

package main

import (
	"bufio"
	"os"
	"strconv"
)

func binSearch(arr *[]int, elem int) bool {
	l, r := 0, len(*arr)-1
	for l < r {
		m := (l + r) / 2
		if (*arr)[m] >= elem {
			r = m
		} else {
			l = m + 1
		}

	}
	return (*arr)[l] == elem

}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, 0, n)
	for i := 0; i < n && scanner.Scan(); i++ {
		num, _ := strconv.Atoi(scanner.Text())
		arr = append(arr, num)
	}

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for i := 0; i < k && scanner.Scan(); i++ {
		num, _ := strconv.Atoi(scanner.Text())
		if binSearch(&arr, num) {
			writer.WriteString("YES\n")
		} else {
			writer.WriteString("NO\n")
		}

	}

}
