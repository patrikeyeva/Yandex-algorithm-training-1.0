//https://contest.yandex.ru/contest/27472/problems/G/

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
	var nums []int

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		if len(nums) < 3 {
			nums = append(nums, num)
			numsSort(&nums)
			continue
		}
		if abs(num) < abs(nums[0]) {
			continue
		}
		nums[0] = num
		numsSort(&nums)

	}
	sort.Ints(nums)

	n := len(nums)

	if nums[0]*nums[1] > nums[n-1]*nums[n-2] {
		fmt.Printf("%d %d", nums[0], nums[1])
	} else {
		fmt.Printf("%d %d", nums[n-2], nums[n-1])
	}

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func numsSort(arr *[]int) {
	n := len(*arr)
	for i := 1; i <= n-1; i++ {
		if abs((*arr)[i-1]) <= abs((*arr)[i]) {
			break
		}
		(*arr)[i], (*arr)[i-1] = (*arr)[i-1], (*arr)[i]

	}
}
