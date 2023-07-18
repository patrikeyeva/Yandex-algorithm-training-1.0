//https://contest.yandex.ru/contest/27665/problems/B/

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

	dict := make(map[string]int)

	for scanner.Scan() {
		word := scanner.Text()
		if count, ok := dict[word]; ok {
			writer.WriteString(strconv.Itoa(count))
			writer.WriteString(" ")

		} else {
			writer.WriteString("0 ")

		}
		dict[word]++

	}

}
