//https://contest.yandex.ru/contest/27665/problems/E/

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

	blocks := make(map[int]int)

	for i := 0; i < n && scanner.Scan(); i++ {
		wh := strings.Split(scanner.Text(), " ")
		w, _ := strconv.Atoi(wh[0])
		h, _ := strconv.Atoi(wh[1])

		if elem, ok := blocks[w]; ok && h > elem || !ok {
			blocks[w] = h
		}
	}

	sum := 0
	for _, elem := range blocks {
		sum += elem
	}
	fmt.Print(sum)

}
