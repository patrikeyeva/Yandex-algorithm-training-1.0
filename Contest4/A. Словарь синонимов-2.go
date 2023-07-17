//https://contest.yandex.ru/contest/27665/problems/A/

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

	list := make([]string, 0, n*2)

	for i := 0; i < n && scanner.Scan(); i++ {
		words := strings.Fields(scanner.Text())
		list = append(list, words...)
	}

	scanner.Scan()
	word := scanner.Text()

	for i, elem := range list {
		if elem == word {
			if i%2 == 0 {
				fmt.Print(list[i+1])
			} else {
				fmt.Print(list[i-1])
			}
			break
		}
	}
}
