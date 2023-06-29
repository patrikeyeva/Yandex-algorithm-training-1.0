//https://contest.yandex.ru/contest/27472/problems/F/

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

	arr := make([]uint8, n)

	scanner.Scan()
	valuesStr := strings.Fields(scanner.Text())
	for i, elem := range valuesStr {
		value, _ := strconv.Atoi(elem)
		arr[i] = uint8(value)
	}

	add := make([]uint8, n)
	count := 0
	last_idx := 0
	i, j := 0, n-1

	for i < j {

		if arr[i] == arr[j] {
			i++
			j--
			continue
		}

		for ; last_idx < i; last_idx++ {
			add[count] = arr[last_idx]
			count++
		}
		j = n - 1
		last_idx = i + 1
		add[count] = arr[i]
		count++
		i++
	}

	fmt.Println(count)
	for ; count > 0; count-- {
		fmt.Printf("%d ", add[count-1])
	}

}
