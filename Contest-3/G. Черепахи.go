//https://contest.yandex.ru/contest/27663/problems/G/

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
	n, _ := strconv.Atoi(scanner.Text())
	count := n
	set := make(map[int]bool)

	for i := 0; i < n && scanner.Scan(); i++ {
		values := strings.Fields(scanner.Text())
		a, _ := strconv.Atoi(values[0])
		b, _ := strconv.Atoi(values[1])
		if a+b != n-1 || a < 0 || b < 0 {
			count--
		} else if _, ok := set[a]; ok {
			count--
		} else {
			set[a] = true
		}
	}
	fmt.Print(count)
}
