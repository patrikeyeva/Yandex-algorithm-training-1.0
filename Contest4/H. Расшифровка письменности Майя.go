//https://contest.yandex.ru/contest/27665/problems/H/

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
	const bufferSize = 3000002
	buffer := make([]byte, bufferSize)
	scanner.Buffer(buffer, bufferSize)

	scanner.Scan()
	nums := strings.Split(scanner.Text(), " ")
	g, _ := strconv.Atoi(nums[0])
	s, _ := strconv.Atoi(nums[1])

	scanner.Scan()
	W := scanner.Text()

	scanner.Scan()
	S := scanner.Text()

	pattern := make(map[byte]int)
	for idx := range W {
		pattern[W[idx]]++
	}
	ans := 0
	check := 0
	word := make(map[byte]int)
	left, right := 0, g

	for i := left; i < right; i++ {
		char := S[i]
		correctCount, ok := pattern[char]
		if !ok { // нет такой буквы в шаблоне
			check--
			continue
		}
		word[char]++
		if correctCount == word[char] {
			check++
		} else if correctCount == word[char]-1 {
			check--
		}

	}
	if check == len(pattern) {
		ans++
	}
	left++
	right++

	for right <= s {

		pred := S[left-1]
		correctCount, ok := pattern[pred]
		if !ok {
			check++
		} else {
			word[pred]--
			if correctCount == word[pred] {
				check++
			} else if correctCount == word[pred]+1 {
				check--
			}
		}

		curr := S[right-1]
		correctCount, ok = pattern[curr]
		if !ok { // нет такой буквы в шаблоне
			check--
		} else {
			word[curr]++
			if correctCount == word[curr] {
				check++
			} else if correctCount == word[curr]-1 {
				check--
			}
		}

		if check == len(pattern) {
			ans++
		}
		left++
		right++
	}

	fmt.Print(ans)

}
