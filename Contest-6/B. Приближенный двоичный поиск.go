//https://contest.yandex.ru/contest/27844/problems/B/

package main

import (
	"bufio"
	"os"
	"strconv"
)

func leftBinSearch(arr *[]int, elem int) (int, int) {
	l, r := 0, len(*arr)-1
	for l < r {
		m := (l + r) / 2
		if (*arr)[m] >= elem {
			r = m
		} else {
			l = m + 1
		}

	}
	return (*arr)[l], l

}

func rightBinSearch(arr *[]int, elem, right int) int {
	l, r := 0, right
	for l < r {
		m := (l + r + 1) / 2
		if (*arr)[m] <= elem {
			l = m
		} else {
			r = m - 1
		}

	}
	return (*arr)[l]

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
		lastElem, right := leftBinSearch(&arr, num)
		firstElem := rightBinSearch(&arr, num, right)

		if num-firstElem <= lastElem-num {
			writer.WriteString(strconv.Itoa(firstElem))
			writer.WriteByte('\n')

		} else {
			writer.WriteString(strconv.Itoa(lastElem))
			writer.WriteByte('\n')

		}

	}

}

