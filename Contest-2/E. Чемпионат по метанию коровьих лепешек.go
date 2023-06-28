//https://contest.yandex.ru/contest/27472/problems/E/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Tuple struct {
	value uint16
	idx   int
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// Read the number of participants
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arr := make([]uint16, n)
	sort_arr := make([]Tuple, n)

	for i := 0; i < n && scanner.Scan(); i++ {
		num, _ := strconv.Atoi(scanner.Text())
		sort_arr[i].value = uint16(num)
		sort_arr[i].idx = i
		arr[i] = sort_arr[i].value
	}

	sort.SliceStable(sort_arr, func(i, j int) bool {
		return sort_arr[i].value > sort_arr[j].value
	})

	max_elem := sort_arr[0].value
	min_idx := sort_arr[0].idx

	for _, elem := range sort_arr {
		if elem.value != max_elem {
			break
		}
		if elem.idx < min_idx {
			min_idx = elem.idx
		}
	}

	res := 0
	pred := sort_arr[0].value
	rate := 0

	for index, elem := range sort_arr {
		check := (pred == elem.value)
		pred = elem.value
		if !check {
			rate = index
		}
		if elem.idx <= min_idx {
			continue
		}
		if elem.value%10 != 5 {
			continue
		}
		if elem.idx == n-1 || arr[elem.idx+1] >= elem.value {
			continue
		}
		res = rate + 1
		break
	}

	fmt.Print(res)
}
