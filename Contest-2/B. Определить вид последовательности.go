//https://contest.yandex.ru/contest/27472/problems/B/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var elem, pred_elem int
	var CONSTANT, ASCENDING, WEAKLY_ASCENDING, DESCENDING, WEAKLY_DESCENDING bool = true, true, true, true, true

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	if scanner.Scan() {
		pred_elem, _ = strconv.Atoi(scanner.Text())
	}
	if scanner.Scan() {
		elem, _ = strconv.Atoi(scanner.Text())
	}

	for elem != -2000000000 && scanner.Scan() {

		CONSTANT = CONSTANT && pred_elem == elem
		ASCENDING = ASCENDING && pred_elem < elem
		WEAKLY_ASCENDING = WEAKLY_ASCENDING && pred_elem <= elem
		DESCENDING = DESCENDING && pred_elem > elem
		WEAKLY_DESCENDING = WEAKLY_DESCENDING && pred_elem >= elem

		pred_elem = elem
		elem, _ = strconv.Atoi(scanner.Text())
	}

	switch {
	case CONSTANT:
		fmt.Print("CONSTANT")
	case ASCENDING:
		fmt.Print("ASCENDING")
	case WEAKLY_ASCENDING:
		fmt.Print("WEAKLY ASCENDING")
	case DESCENDING:
		fmt.Print("DESCENDING")
	case WEAKLY_DESCENDING:
		fmt.Print("WEAKLY DESCENDING")
	default:
		fmt.Print("RANDOM")
	}
}
