//https://contest.yandex.ru/contest/27472/problems/A/

package main

import (
	"fmt"
	"os"
)

func main() {

	check := true
	var elem, pred_elem int

	fmt.Scan(&pred_elem)
	for {
		if _, err := fmt.Fscan(os.Stdin, &elem); err != nil {
			break
		}

		if check && elem <= pred_elem {
			check = false
		}
		pred_elem = elem

	}

	if check {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}

}
