//https://contest.yandex.ru/contest/27663/problems/A/

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

	set := make(map[int]struct{})

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		set[num] = struct{}{}
	}

	fmt.Print(len(set))
}


/*
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

	set := make(map[int]bool)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		set[num] = true

	}

	fmt.Print(len(set))
}


*/
