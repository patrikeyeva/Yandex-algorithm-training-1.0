package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var max, min []int

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())

		if len(max) < 3 {
			max = append(max, num)
			numsSort(&max)
		} else if num > max[2] {
			max[2] = num
			numsSort(&max)
		}

		if len(min) < 2 {
			min = append(min, num)
			sort.Ints(min)
		} else if num < min[1] {
			min[1] = num
			sort.Ints(min)

		}

	}


	if max[0]*max[1]*max[2] > max[0]*min[0]*min[1] {
		fmt.Printf("%d %d %d", max[0], max[1], max[2])
	} else {
		fmt.Printf("%d %d %d", max[0], min[0], min[1])
	}

}

func numsSort(arr *[]int) {
	n := len(*arr)
	for i := n - 1; i > 0; i-- {
		if (*arr)[i-1] >= (*arr)[i] {
			break
		}
		(*arr)[i], (*arr)[i-1] = (*arr)[i-1], (*arr)[i]

	}
}
