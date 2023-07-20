//https://contest.yandex.ru/contest/27665/problems/I/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	dict := make(map[string]bool)
	dictCheck := make(map[string]bool)

	for i := 0; i < n && scanner.Scan(); i++ {
		word := scanner.Text()
		dict[word] = true
		dictCheck[strings.ToLower(word)] = true
	}

	errors := 0
	for scanner.Scan() {
		word := scanner.Text()

		existsInDict := dict[word]
		existsInDictCheck := dictCheck[strings.ToLower(word)]

		if !existsInDict && existsInDictCheck { //слово есть в словаре, но в некоректном регистре
			errors++
			continue
		}

		// слова нет в словаре
		upperFound := false
		for _, char := range word {
			if unicode.IsUpper(char) {
				if upperFound {
					errors++
					break
				}
				upperFound = true
			}
		}
		if !upperFound {
			errors++
		}
	}

	fmt.Print(errors)
}
