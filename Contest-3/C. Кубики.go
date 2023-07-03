//https://contest.yandex.ru/contest/27663/problems/C/


import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func intersection(anya, borya *map[uint]bool, res *[]uint) {

	for num := range *anya {
		if (*borya)[num] {
			*res = append(*res, num)
			delete(*anya, num)
			delete(*borya, num)
		}
	}
}

func out(res *[]uint, writer *bufio.Writer) {
	writer.WriteString(strconv.Itoa(len(*res)) + "\n")
	sort.Slice(*res, func(i, j int) bool {
		return (*res)[i] < (*res)[j]

	})
	for i := 0; i < len(*res); i++ {
		if i > 0 {
			writer.WriteString(" ")

		}
		writer.WriteString(strconv.FormatUint(uint64((*res)[i]), 10))
	}
	writer.WriteString("\n")

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	anya := make(map[uint]bool)
	borya := make(map[uint]bool)

	for i := 0; i < n && scanner.Scan(); i++ {
		num, _ := strconv.Atoi(scanner.Text())
		anya[uint(num)] = true
	}
	for i := 0; i < m && scanner.Scan(); i++ {
		num, _ := strconv.Atoi(scanner.Text())
		borya[uint(num)] = true
	}

	res := make([]uint, 0)
	intersection(&anya, &borya, &res)

	arr1 := make([]uint, 0, len(anya))
	arr2 := make([]uint, 0, len(borya))
	for k := range anya {
		arr1 = append(arr1, k)
	}

	for k := range borya {
		arr2 = append(arr2, k)
	}

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	out(&res, writer)
	out(&arr1, writer)
	out(&arr2, writer)
}
