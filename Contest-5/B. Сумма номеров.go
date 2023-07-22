//https://contest.yandex.ru/contest/27794/problems/B/

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
	K, _ := strconv.Atoi(scanner.Text())

	prefixSum := make([]int, 1, n+1)
	for i := 0; i < n && scanner.Scan(); i++ {
		num, _ := strconv.Atoi(scanner.Text())
		newElem := prefixSum[i] + num
		prefixSum = append(prefixSum, newElem)
	}

	left, right := 0, 1
	count := 0

	for right < n + 1 {
		sum := prefixSum[left] + K
		if sum == prefixSum[right] {
			count++
			left++
		} else if sum < prefixSum[right] {
			left++
		} else {
			right++
		}
	}

	fmt.Print(count)

}
