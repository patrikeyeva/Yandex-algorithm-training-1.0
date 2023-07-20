//https://contest.yandex.ru/contest/27665/problems/G/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	bank := make(map[string]int)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for scanner.Scan() {
		str := strings.Split(scanner.Text(), " ")

		switch str[0] {
		case "DEPOSIT":
			name := str[1]
			sum, _ := strconv.Atoi(str[2])
			bank[name] += sum

		case "WITHDRAW":
			name := str[1]
			sum, _ := strconv.Atoi(str[2])
			bank[name] -= sum

		case "BALANCE":
			name := str[1]
			if balance, ok := bank[name]; ok {
				writer.WriteString(strconv.Itoa(balance))
				writer.WriteString("\n")
				//fmt.Printf("%d\n", balance)
			} else {
				writer.WriteString("ERROR\n")
			}

		case "TRANSFER":
			name1, name2 := str[1], str[2]
			sum, _ := strconv.Atoi(str[3])
			bank[name1] -= sum
			bank[name2] += sum

		case "INCOME":
			p, _ := strconv.ParseFloat(str[1], 64)
			for key, value := range bank {
				if value > 0 {
					res := float64(value) * (1.0 + p/100)
					bank[key] = int(res)
				}
			}
		}
	}
}
