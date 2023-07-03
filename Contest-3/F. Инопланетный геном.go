//https://contest.yandex.ru/contest/27663/problems/F/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func readDate(gen1 *[]byte, reader *bufio.Reader) {
	for {
		line, ispref, _ := reader.ReadLine()
		*gen1 = append(*gen1, line...)
		if !ispref {
			break
		}
	}

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var gen1, gen2 []byte

	readDate(&gen1, reader)
	readDate(&gen2, reader)

	counter := make(map[string]uint)

	for i := 0; i < len(gen1)-1; i++ {

		key := string(gen1[i]) + string(gen1[i+1])
		counter[key]++

	}

	var ans uint = 0

	for i := 0; i < len(gen2)-1; i++ {

		key := string(gen2[i]) + string(gen2[i+1])
		if elem, ok := counter[key]; ok {
			ans += elem
			delete(counter, key)
		}

	}

	fmt.Print(ans)

}
