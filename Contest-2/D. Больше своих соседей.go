//https://contest.yandex.ru/contest/27472/problems/D/

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
	pred, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	curr, _ := strconv.Atoi(scanner.Text())

	check := false
	count := 0

	if curr > pred {
		check = true
	}

	for scanner.Scan() {
		next, _ := strconv.Atoi(scanner.Text())
		if curr > next && check {
			count++
			check = false
		}
		if curr < next {
			check = true
		}
		curr = next

	}
	fmt.Print(count)

}
