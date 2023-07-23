//https://contest.yandex.ru/contest/27794/problems/F/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	conditioner := make([]int, 1001)
	for i := 0; i < n && scanner.Scan(); i++ {
		power, _ := strconv.Atoi(scanner.Text())
		conditioner[power]++
	}

	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	models := make([]int, 1001)
	for i := 0; i < m && scanner.Scan(); i++ {
		power, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		price, _ := strconv.Atoi(scanner.Text())
		if power > models[price] {
			models[price] = power
		}
	}

	sum := 0
	ptrCond, ptrModel := 0, 0

	for ptrCond < len(conditioner) {
		if conditioner[ptrCond] == 0 {
			ptrCond++
		} else {
			needPower, needNumber := ptrCond, conditioner[ptrCond]
			for ptrModel < len(models) {
				if models[ptrModel] < needPower {
					ptrModel++
				} else {
					price := ptrModel
					sum += price * needNumber
					ptrCond++
					break
				}
			}
		}

	}
	fmt.Print(sum)
}
