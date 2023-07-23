//https://contest.yandex.ru/contest/27794/problems/E/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	treeColor := make([]int, 0, n)
	for i := 0; i < n && scanner.Scan(); i++ {
		num, _ := strconv.Atoi(scanner.Text())
		treeColor = append(treeColor, num)
	}

	counterColor := make([]int, k+1)

	left, right := 0, 1
	minLeft, minRight, minLen := 0, n-1, n

	counterColor[treeColor[0]]++
	check := 1

	for right < n {
		rightElem, leftElem := treeColor[right], treeColor[left]

		if leftElem == rightElem {
			left++
			for leftElem = treeColor[left]; counterColor[leftElem] > 1; {
				counterColor[leftElem]--
				left++
				leftElem = treeColor[left]
			}
		} else {
			if counterColor[rightElem] == 0 {
				check++
			}
			counterColor[rightElem]++
		}

		if check == k && right-left+1 < minLen {
			minLeft, minRight, minLen = left, right, right-left+1
		}
		right++

	}

	fmt.Printf("%d %d", minLeft+1, minRight+1)

}
