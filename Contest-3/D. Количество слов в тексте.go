//https://contest.yandex.ru/contest/27663/problems/D/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	set := make(map[string]bool)

	for scanner.Scan() {
		word := scanner.Text()
		set[word] = true

	}
	fmt.Print(len(set))
}
