//https://contest.yandex.ru/contest/27663/problems/I/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	counter := make(map[string]uint)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n && scanner.Scan(); i++ {
		m, _ := strconv.Atoi(scanner.Text())
		for j := 0; j < m && scanner.Scan(); j++ {
			lang := scanner.Text()
			counter[lang]++
		}

	}

	all := make([]string, 0)
	once := make([]string, 0)

	for key, value := range counter {
		if value == uint(n) {
			all = append(all, key)
		}
		if value > 0 {
			once = append(once, key)

		}

	}

	fmt.Println(len(all))
	for _, elem := range all {
		fmt.Println(elem)
	}

	fmt.Println(len(once))
	for _, elem := range once {
		fmt.Println(elem)
	}
}
