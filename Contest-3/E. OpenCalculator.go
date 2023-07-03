//https://contest.yandex.ru/contest/27663/problems/E/


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
	scanner.Split(bufio.ScanWords)
	set := make([]bool, 10)
	res := 0

	for i := 0; i < 3 && scanner.Scan(); i++ {
		num, _ := strconv.Atoi(scanner.Text())
		set[num] = true

	}
	scanner.Scan()
	str := scanner.Text()
	nums := strings.Split(str, "")

	for _, elem := range nums {
		num, _ := strconv.Atoi(elem)
		if !set[num] {
			res++
			set[num] = true
		}
	}

	fmt.Print(res)

}

