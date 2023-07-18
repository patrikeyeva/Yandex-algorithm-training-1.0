//https://contest.yandex.ru/contest/27665/problems/C/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	dict := make(map[string]uint)
	var ans string
	var max_count, count uint

	for scanner.Scan() {
		word := scanner.Text()
		dict[word]++
		count = dict[word]

		if count > max_count {
			max_count = count
			ans = word
		} else if count == max_count && word < ans {
			ans = word
		}

	}
	fmt.Print(ans)
}
