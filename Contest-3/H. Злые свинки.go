//https://contest.yandex.ru/contest/27663/problems/H/

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
	counter := make(map[uint]uint)

	for i := 0; i < n && scanner.Scan(); i++ {
		values := strings.Fields(scanner.Text())
		x, _ := strconv.Atoi(values[0])
		counter[uint(x)]++
	}
	fmt.Print(len(counter))

}
