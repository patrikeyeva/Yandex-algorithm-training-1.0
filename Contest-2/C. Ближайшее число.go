//https://contest.yandex.ru/contest/27472/problems/C/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(a, b int) int {
	if a-b < 0 {
		return -(a - b)
	}
	return a - b
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, n)
	scanner.Scan()
	valuesStr := strings.Fields(scanner.Text())

	for i := range valuesStr {
		arr[i], _ = strconv.Atoi(valuesStr[i])
	}
	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())

	ans, diff := arr[0], abs(value, arr[0])

	for _, elem := range arr {

		res := abs(value, elem)

		if res < diff {
			diff = res
			ans = elem
		}

	}

	fmt.Print(ans)

}
