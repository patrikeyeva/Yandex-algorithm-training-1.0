//https://contest.yandex.ru/contest/27665/problems/D/

package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	clicks := make([]int, 0, n)
	for i := 0; i < n && scanner.Scan(); i++ {
		num, _ := strconv.Atoi(scanner.Text())
		clicks = append(clicks, num)
	}

	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < k && scanner.Scan(); i++ {
		idx, _ := strconv.Atoi(scanner.Text())
		idx--
		clicks[idx]--

	}

	for _, elem := range clicks {
		if elem >= 0 {
			writer.WriteString("NO\n")
		} else {
			writer.WriteString("YES\n")
		}

	}

}
