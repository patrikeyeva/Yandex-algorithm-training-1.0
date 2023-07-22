//https://contest.yandex.ru/contest/27794/problems/A/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	shirts := make([]int, 0, n)
	for i := 0; i < n && scanner.Scan(); i++ {
		num, _ := strconv.Atoi(scanner.Text())
		shirts = append(shirts, num)
	}

	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	pants := make([]int, 0, m)
	for i := 0; i < m && scanner.Scan(); i++ {
		num, _ := strconv.Atoi(scanner.Text())
		pants = append(pants, num)
	}

	minDiff := abs(shirts[0] - pants[0])
	shirtId, pantsId := 0, 0

	sPtr, pPtr := 0, 0

	for sPtr != len(shirts)-1 || pPtr != len(pants)-1 {

		if pPtr == len(pants)-1 || (shirts[sPtr] < pants[pPtr] && sPtr != len(shirts)-1) {
			sPtr = min(sPtr+1, len(shirts)-1)
		} else {
			pPtr = min(pPtr+1, len(pants)-1)
		}
		diff := abs(shirts[sPtr] - pants[pPtr])
		if diff < minDiff {
			minDiff, shirtId, pantsId = diff, sPtr, pPtr
		}

	}

	fmt.Printf("%d %d", shirts[shirtId], pants[pantsId])

}
