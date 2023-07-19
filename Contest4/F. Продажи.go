https://contest.yandex.ru/contest/27665/problems/F/


package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseKey(key *string) (string, string) {
	parts := strings.Split(*key, " ")
	return parts[0], parts[1]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	info := make(map[string]int)

	for scanner.Scan() {
		str := scanner.Text()
		idx := strings.LastIndex(str, " ")
		key := str[:idx]
		count, _ := strconv.Atoi(str[idx+1:])
		info[key] += count
	}

	keys := make([]string, 0, len(info))
	for key := range info {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	currentName := ""
	for _, key := range keys {
		name, product := parseKey(&key)
		if name != currentName {
			writer.WriteString(name + ":\n")
			currentName = name
		}
		count := strconv.Itoa(info[key])
		writer.WriteString(product + " " + count + "\n")
	}
}
