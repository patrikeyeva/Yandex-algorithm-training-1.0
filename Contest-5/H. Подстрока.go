//https://contest.yandex.ru/contest/27794/problems/H/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	bufferSize := 100002
	scanner.Buffer(make([]byte, bufferSize), bufferSize)

	scanner.Scan()
	nums := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(nums[0])
	k, _ := strconv.Atoi(nums[1])

	scanner.Scan()
	str := scanner.Text()

	counter := [26]int{}

	maxLeft, maxLen := 0, 1

	for left, right := 0, 0; right < n; right++ {
		newElem := str[right] - 97
		counter[newElem]++

		for counter[newElem] > k {
			counter[str[left]-97]--
			left++
		}

		if right-left+1 > maxLen {
			maxLeft, maxLen = left, right-left+1
		}
	}

	fmt.Printf("%d %d", maxLen, maxLeft+1)
}
