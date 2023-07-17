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

	dict := make(map[string]string)

	for i := 0; i < n && scanner.Scan(); i++ {
		words := strings.Fields(scanner.Text())
		dict[words[0]] = words[1]
		dict[words[1]] = words[0]
	}

	scanner.Scan()
	word := scanner.Text()
	fmt.Print(dict[word])
}
