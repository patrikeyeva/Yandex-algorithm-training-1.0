//https://contest.yandex.ru/contest/27665/problems/J/

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

type Tuple struct {
	count, idx int
}

func main() {
	var n int
	var C, D string
	fmt.Scan(&n, &C, &D)
	keywords := make(map[string]bool)
	scanner := bufio.NewScanner(os.Stdin)

	caseSensitive := (C == "yes")
	startNum := (D == "yes")

	for i := 0; i < n && scanner.Scan(); i++ {
		word := scanner.Text()
		if !caseSensitive {
			word = strings.ToLower(word)
		}
		keywords[word] = true
	}

	regexpStr := `[a-zA-Z0-9_]*[a-zA-Z_]+[a-zA-Z0-9_]*`
	regexPattern := regexp.MustCompile(regexpStr)

	idCount := make(map[string]Tuple)
	for scanner.Scan() {
		line := scanner.Text()
		matches := regexPattern.FindAllString(line, -1)

		for _, elem := range matches {
			if !startNum && unicode.IsDigit(rune(elem[0])) {
				continue
			}
			if !caseSensitive { // C == no - нечувствительны
				elem = strings.ToLower(elem)
			}
			if !keywords[elem] { // не ключевое слово
				value, ok := idCount[elem]
				if !ok {
					idCount[elem] = Tuple{0, len(idCount)}
				} else {
					value.count++
					idCount[elem] = value
				}
			}
		}

	}
	ans := ""
	maxCount := 0
	ansId := len(idCount) + 2
	for key, value := range idCount {
		count := value.count
		id := value.idx
		if count > maxCount || (count == maxCount && id < ansId) {
			ans, maxCount, ansId = key, count, id
		}
	}
	fmt.Print(ans)
}
