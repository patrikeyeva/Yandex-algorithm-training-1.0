//https://contest.yandex.ru/contest/27663/problems/B/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func strToInt(str *[]string, set *map[int]struct{}) {
	for _, elem := range *str {
		num, _ := strconv.Atoi(elem)
		(*set)[num] = struct{}{}
	}
}

func intersection(set1, set2 *map[int]struct{}, res *[]int) {

	for num := range *set1 {
		if _, ok := (*set2)[num]; ok {
			*res = append(*res, num)
		}

	}

}

func main() {

	reader := bufio.NewReader(os.Stdin)

	line1, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print("error")
	}
	line2, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print("error")
	}
	arr1 := strings.Fields(line1)
	arr2 := strings.Fields(line2)

	set1 := make(map[int]struct{}, len(arr1))
	set2 := make(map[int]struct{}, len(arr2))

	strToInt(&arr1, &set1)
	strToInt(&arr2, &set2)

	res := make([]int, 0)

	intersection(&set1, &set2, &res)

	sort.Ints(res)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for i, elem := range res {
		if i > 0 {
			writer.WriteString(" ")
		}
		writer.WriteString(strconv.Itoa(elem))
	}
}
