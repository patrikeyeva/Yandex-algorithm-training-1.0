//https://contest.yandex.ru/contest/27794/problems/C/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arrY := make([]int, 0, n)
	prefixSumLeft := make([]int, 1, n)

	scanner.Scan()
	nums := strings.Fields(scanner.Text())
	y, _ := strconv.Atoi(nums[1])
	arrY = append(arrY, y)

	for i := 1; i < n && scanner.Scan(); i++ {
		nums := strings.Fields(scanner.Text())
		y, _ := strconv.Atoi(nums[1])
		predY := arrY[i-1]
		arrY = append(arrY, y)

		newElem := prefixSumLeft[i-1] + max(0, y-predY)
		prefixSumLeft = append(prefixSumLeft, newElem)

	}

	prefixSumRight := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		y, predY := arrY[i], arrY[i+1]
		prefixSumRight[i] = prefixSumRight[i+1] + max(0, y-predY)
	}

	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	var sum int

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for i := 0; i < m && scanner.Scan(); i++ {
		nums := strings.Fields(scanner.Text())
		left, _ := strconv.Atoi(nums[0])
		right, _ := strconv.Atoi(nums[1])
		left--
		right--
		if left < right {
			sum = prefixSumLeft[right] - prefixSumLeft[left]
		} else {
			sum = prefixSumRight[right] - prefixSumRight[left]

		}
		writer.WriteString(strconv.Itoa(sum))
		writer.WriteByte('\n')

	}

}
