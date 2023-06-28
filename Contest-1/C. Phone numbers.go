/*
https://contest.yandex.ru/contest/27393/problems/C/
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func processStr(str *string) {

	*str = strings.ReplaceAll(*str, "(", "")
	*str = strings.ReplaceAll(*str, ")", "")
	*str = strings.ReplaceAll(*str, "-", "")

	if len(*str) == 7 {
		*str = "495" + (*str)
	}
}

func checkNumber(want_add, num *string) bool {

	return (*want_add)[len(*want_add)-10:] == (*num)[len(*num)-10:]

}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numAdd := scanner.Text()
	processStr(&numAdd)

	nums := make([]string, 3)
	for i := range nums {
		scanner.Scan()
		nums[i] = scanner.Text()
		processStr(&nums[i])
	}

	for _, number := range nums {
		if checkNumber(&numAdd, &number) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
