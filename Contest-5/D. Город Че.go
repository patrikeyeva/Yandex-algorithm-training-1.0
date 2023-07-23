//https://contest.yandex.ru/contest/27794/problems/D/


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

	scanner.Scan()
	r, _ := strconv.Atoi(scanner.Text())

	distance := make([]int, 0, n)
	for i := 0; i < n && scanner.Scan(); i++ {
		num, _ := strconv.Atoi(scanner.Text())
		distance = append(distance, num)
	}

	left, right := 0, 1
	counter := 0
	for right < n {
		diff := distance[right] - distance[left]
		if diff > r {
			counter += n - right
			left++
		} else {
			right++
		}
	}

	fmt.Print(counter)

}
