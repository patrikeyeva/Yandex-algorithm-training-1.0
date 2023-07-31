//https://contest.yandex.ru/contest/27844/problems/H/


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

	scanner.Scan()
	nums := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(nums[0])
	k, _ := strconv.Atoi(nums[1])

	var sumLength uint64 = 0
	segments := make([]int, 0, n)
	for i := 0; i < n && scanner.Scan(); i++ {
		elem, _ := strconv.Atoi(scanner.Text())
		segments = append(segments, elem)
		sumLength += uint64(elem)
	}

	l, r := uint64(0), sumLength

	for l < r {
		len := (l + r + 1) / 2

		var numberOfSegments uint64 = 0
		for _, segment := range segments {
			numberOfSegments += uint64(segment) / len

		}

		if numberOfSegments >= uint64(k) {
			l = len
		} else {
			r = len - 1

		}

	}

	fmt.Print(l)
}
